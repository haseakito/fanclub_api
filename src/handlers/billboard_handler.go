package handlers

import (
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/hackgame-org/fanclub_api/ent"
	"github.com/hackgame-org/fanclub_api/ent/billboard"
	"github.com/hackgame-org/fanclub_api/requests"
	"github.com/labstack/echo/v4"
)

type BillboardHandler struct {
	db  *ent.Client
	cld *cloudinary.Cloudinary
}

func NewBillboardHandler(db *ent.Client, cld *cloudinary.Cloudinary) *BillboardHandler {
	return &BillboardHandler{
		db:  db,
		cld: cld,
	}
}

func (h BillboardHandler) CreateBillboard(c echo.Context) error {
	// Bind the request data to PostRequet
	var req requests.BillboardRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Create a new billboard
	res, err := h.db.Billboard.
		Create().
		SetTitle(req.Title).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h BillboardHandler) UploadFile(c echo.Context) error {
	// Get billboard id from request
	billboardID := c.Param("id")

	// Parse billboard ID string into UUID
	billboardUUID, err := uuid.Parse(billboardID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Query the billboard with billboard id
	bill, err := h.db.Billboard.
		Query().
		WithAsset().
		Where(billboard.ID(billboardUUID)).
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// If there is already an asset associated with this billboard, delete it first before uploading a new file˝
	if bill.Edges.Asset != nil {
		// Delete the object from cloudinary
		_, err = h.cld.Upload.Destroy(c.Request().Context(), uploader.DestroyParams{
			PublicID:     bill.Edges.Asset.PublicID,
			ResourceType: bill.Edges.Asset.ResourceType,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	// Retrieve file from request form
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request file	
	if err := requests.ValidateFile(file); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())		
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer src.Close()

	// Upload an asset to cloudinary
	resp, err := h.cld.Upload.Upload(
		c.Request().Context(),
		src,
		uploader.UploadParams{},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Start a new transaction
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer tx.Rollback()

	// Insert a new asset
	asset, err := tx.Asset.
		Create().
		SetPublicID(resp.PublicID).
		SetURL(resp.SecureURL).
		SetResourceType(resp.ResourceType).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Update a billboard with the asset
	bill, err = tx.Billboard.
		UpdateOneID(billboardUUID).
		SetAsset(asset).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, bill)
}

func (h BillboardHandler) GetBillboards(c echo.Context) error {
	// Query billboards
	bills, err := h.db.Billboard.
		Query().
		WithAsset().
		All(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, bills)
}

func (h BillboardHandler) GetBillboard(c echo.Context) error {
	// Get billboard id from request
	billboardID := c.Param("id")

	// Parse billboard ID string into UUID
	billboardUUID, err := uuid.Parse(billboardID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Query the billboard with billboard id
	bill, err := h.db.Billboard.
		Query().
		WithAsset().
		Where(billboard.ID(billboardUUID)).
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, bill)
}

func (h BillboardHandler) UpdateBillboard(c echo.Context) error {
	// Get billboard id from request
	billboardID := c.Param("id")

	// Parse billboard ID string into UUID
	billboardUUID, err := uuid.Parse(billboardID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Bind the request data to PostRequet
	var req requests.BillboardRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate request data
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Update the fields
	bill, err := h.db.Billboard.
		UpdateOneID(billboardUUID).
		SetTitle(req.Title).
		SetDescription(req.Description).
		Save(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, bill)
}

func (h BillboardHandler) DeleteBillboard(c echo.Context) error {
	// Get billboard id from request
	billboardID := c.Param("id")

	// Parse billboard ID string into UUID
	billboardUUID, err := uuid.Parse(billboardID)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Start a new transaction
	tx, err := h.db.Tx(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer tx.Rollback()

	// Query the billboard with billboard id
	bill, err := tx.Billboard.
		Query().
		WithAsset().
		Where(billboard.ID(billboardUUID)).
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}

	// Delete the object from cloudinary
	if bill.Edges.Asset != nil {
		_, err = h.cld.Upload.Destroy(c.Request().Context(), uploader.DestroyParams{
			PublicID:     bill.Edges.Asset.PublicID,
			ResourceType: bill.Edges.Asset.ResourceType,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	// Delete the billboard with billboard id
	if err := tx.Billboard.
		DeleteOneID(billboardUUID).
		Exec(c.Request().Context()); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Successfully deleted the billboard")
}
