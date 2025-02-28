package controller

import (
	"example-api/config"
	"example-api/model"
	"example-api/model/base"
	"example-api/util"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRole menambahkan role baru ke database dengan validasi nama unik
func CreateRole(c *gin.Context) {

	var role model.Role
	response := base.BaseResponse{
		Status:  util.Success,
		Message: "Role created successfully",
	}

	// Bind JSON & Validasi
	if !util.BindJSONGeneric(c, &role) {
		return
	}

	// Cek apakah role dengan nama yang sama sudah ada
	var count int64
	config.DB.Model(&model.Role{}).Where("name = ?", role.Name).Count(&count)
	if count > 0 {
		response.Status = util.Failed
		response.Message = "Role name already exists"
		c.JSON(http.StatusConflict, response)
		return
	}

	// Simpan role ke database
	if err := config.DB.Create(&role).Error; err != nil {
		response.Status = util.Failed
		response.Message = "Failed to create role"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Jika berhasil, kembalikan response dengan data role baru
	response.Data = role
	c.JSON(http.StatusOK, response)
}

func GetRoles(c *gin.Context) {
	var roles []model.Role
	response := base.BaseResponse{
		Status:  util.Success,
		Message: "Success get all roles",
	}

	if err := config.DB.Find(&roles).Error; err != nil {
		response.Status = util.Failed
		response.Message = "Failed get all roles"
	} else {
		response.Data = roles
	}
	c.JSON(http.StatusOK, response)
}

func GetRoleById(c *gin.Context) {
	var role model.Role
	response := base.BaseResponse{
		Status:  util.Success,
		Message: "Success get role by id",
	}

	id := c.Query("id")
	var count int64
	config.DB.Model(&model.Role{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		response.Status = util.Failed
		response.Message = "Role not found"
		c.JSON(http.StatusNotFound, response)
		return

	}
	config.DB.First(&role, id)
	response.Data = role
	c.JSON(http.StatusOK, response)
}

func UpdateRole(c *gin.Context) {
	var role model.Role
	response := base.BaseResponse{
		Status:  util.Success,
		Message: "Success update role",
	}
	id := c.Query("id")

	var count int64
	config.DB.Model(&model.Role{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		response.Status = util.Failed
		response.Message = "Role not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	if !util.BindJSONGeneric(c, &role) {
		return
	}

	// Cek apakah role dengan nama yang sama sudah ada
	var countName int64
	config.DB.Model(&model.Role{}).Where("name = ? AND id != ?", role.Name, id).Count(&countName)
	if countName > 0 {
		response.Status = util.Failed
		response.Message = "Role name already exists"
		c.JSON(http.StatusConflict, response)
		return
	}

	// Update role di database
	if err := config.DB.Model(&role).Where("id = ?", id).Updates(&role).Error; err != nil {
		response.Status = util.Failed
		response.Message = "Failed to update role"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Data = role
	c.JSON(http.StatusOK, response)
}

func DeleteRole(c *gin.Context) {
	var role model.Role
	response := base.BaseResponse{
		Status:  util.Success,
		Message: "Success delete role",
	}
	id := c.Query("id")

	var count int64
	config.DB.Model(&model.Role{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		response.Status = util.Failed
		response.Message = "Role not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	if err := config.DB.Delete(&role, id).Error; err != nil {
		response.Status = util.Failed
		response.Message = "Failed to delete role"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Data = role
	c.JSON(http.StatusOK, response)
}

// GetRolePage mengambil daftar role dengan pagination
func GetRolePage(c *gin.Context) {
	var roles []model.Role
	var totalRows int64

	// Ambil parameter page dan pageSize dari query string
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// Hitung total jumlah data
	if err := config.DB.Model(&model.Role{}).Count(&totalRows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  util.Failed,
			Message: "Gagal menghitung total data",
		})
		return
	}

	// Hitung total halaman
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))

	// Ambil data dengan pagination
	if err := config.DB.Order("id ASC").Offset(offset).Limit(pageSize).Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  util.Failed,
			Message: "Gagal mengambil data role",
		})
		return
	}

	// Jika tidak ada data di halaman tertentu
	if len(roles) == 0 {
		c.JSON(http.StatusOK, base.BaseResponse{
			Status:  util.Success,
			Message: "Tidak ada data pada halaman ini",
			Data: base.Pagination{
				Page:       page,
				PageSize:   pageSize,
				TotalRows:  totalRows,
				TotalPages: totalPages,
				Data:       []model.Role{},
			},
		})
		return
	}

	// Format response JSON
	c.JSON(http.StatusOK, base.BaseResponse{
		Status:  util.Success,
		Message: "Success get role page",
		Data: base.Pagination{
			Page:       page,
			PageSize:   pageSize,
			TotalRows:  totalRows,
			TotalPages: totalPages,
			Data:       roles,
		},
	})
}
