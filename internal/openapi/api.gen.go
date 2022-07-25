// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// Defines values for LayoutType.
const (
	LayoutTypeALBUM LayoutType = "ALBUM"

	LayoutTypeSQUARE LayoutType = "SQUARE"

	LayoutTypeTIMELINE LayoutType = "TIMELINE"

	LayoutTypeWALL LayoutType = "WALL"
)

// Defines values for TaskType.
const (
	TaskTypeINDEX TaskType = "INDEX"

	TaskTypeLOADCOLOR TaskType = "LOAD_COLOR"

	TaskTypeLOADMETA TaskType = "LOAD_META"
)

// Bounds defines model for Bounds.
type Bounds struct {
	H    *float32 `json:"h,omitempty"`
	True *float32 `json:"true,omitempty"`
	W    *float32 `json:"w,omitempty"`
	X    *float32 `json:"x,omitempty"`
}

// Collection defines model for Collection.
type Collection struct {
	Id CollectionId `json:"id"`

	// Time of latest performed full index
	IndexedAt *time.Time `json:"indexed_at,omitempty"`

	// User-friendly name
	Name *string `json:"name,omitempty"`
}

// CollectionId defines model for CollectionId.
type CollectionId string

// File defines model for File.
type File string

// FileId defines model for FileId.
type FileId int

// ImageHeight defines model for ImageHeight.
type ImageHeight float32

// LayoutType defines model for LayoutType.
type LayoutType string

// Problem defines model for Problem.
type Problem struct {
	// The HTTP status code generated by the origin server for this occurrence of the problem.
	Status *int32 `json:"status,omitempty"`

	// A short summary of the problem type. Written in English and readable for engineers, usually not suited for non technical stakeholders and not localized.
	Title *string `json:"title,omitempty"`
}

// Region defines model for Region.
type Region struct {
	Bounds Bounds      `json:"bounds"`
	Data   *RegionData `json:"data,omitempty"`
	Id     RegionId    `json:"id"`
}

// RegionData defines model for RegionData.
type RegionData map[string]interface{}

// RegionId defines model for RegionId.
type RegionId int

// Scene defines model for Scene.
type Scene struct {
	Bounds    *Bounds `json:"bounds,omitempty"`
	FileCount *int    `json:"file_count,omitempty"`
	Id        SceneId `json:"id"`

	// True while the scene is loading and the dimensions are not yet known.
	Loading *bool `json:"loading,omitempty"`
}

// SceneId defines model for SceneId.
type SceneId string

// SceneParams defines model for SceneParams.
type SceneParams struct {
	CollectionId CollectionId `json:"collection_id"`
	ImageHeight  ImageHeight  `json:"image_height"`
	Layout       LayoutType   `json:"layout"`
	SceneWidth   SceneWidth   `json:"scene_width"`
}

// SceneWidth defines model for SceneWidth.
type SceneWidth float32

// Task defines model for Task.
type Task struct {
	CollectionId *CollectionId `json:"collection_id,omitempty"`

	// Number of items already processed.
	Done *int   `json:"done,omitempty"`
	Id   TaskId `json:"id"`
	Name string `json:"name"`

	// Number of items pending as part of the task.
	Pending *int      `json:"pending,omitempty"`
	Type    *TaskType `json:"type,omitempty"`
}

// TaskId defines model for TaskId.
type TaskId string

// TaskType defines model for TaskType.
type TaskType string

// TileCoord defines model for TileCoord.
type TileCoord int

// FileIdPathParam defines model for FileIdPathParam.
type FileIdPathParam FileId

// FilenamePathParam defines model for FilenamePathParam.
type FilenamePathParam string

// SizePathParam defines model for SizePathParam.
type SizePathParam string

// GetScenesParams defines parameters for GetScenes.
type GetScenesParams struct {
	// Collection ID
	CollectionId CollectionId `json:"collection_id"`
	SceneWidth   *SceneWidth  `json:"scene_width,omitempty"`
	ImageHeight  *ImageHeight `json:"image_height,omitempty"`
	Layout       *LayoutType  `json:"layout,omitempty"`
}

// PostScenesJSONBody defines parameters for PostScenes.
type PostScenesJSONBody SceneParams

// GetScenesSceneIdDatesParams defines parameters for GetScenesSceneIdDates.
type GetScenesSceneIdDatesParams struct {
	Height int `json:"height"`
}

// GetScenesSceneIdRegionsParams defines parameters for GetScenesSceneIdRegions.
type GetScenesSceneIdRegionsParams struct {
	X     float32 `json:"x"`
	Y     float32 `json:"y"`
	W     float32 `json:"w"`
	H     float32 `json:"h"`
	Limit *int    `json:"limit,omitempty"`
}

// GetScenesSceneIdTilesParams defines parameters for GetScenesSceneIdTiles.
type GetScenesSceneIdTilesParams struct {
	TileSize        int       `json:"tile_size"`
	Zoom            int       `json:"zoom"`
	X               TileCoord `json:"x"`
	Y               TileCoord `json:"y"`
	DebugOverdraw   *bool     `json:"debug_overdraw,omitempty"`
	DebugThumbnails *bool     `json:"debug_thumbnails,omitempty"`
}

// GetTasksParams defines parameters for GetTasks.
type GetTasksParams struct {
	// Task type to filter on.
	Type *TaskType `json:"type,omitempty"`

	// Collection ID for the tasks
	CollectionId *CollectionId `json:"collection_id,omitempty"`
}

// PostTasksJSONBody defines parameters for PostTasks.
type PostTasksJSONBody struct {
	CollectionId CollectionId `json:"collection_id"`
	Type         TaskType     `json:"type"`
}

// PostScenesJSONRequestBody defines body for PostScenes for application/json ContentType.
type PostScenesJSONRequestBody PostScenesJSONBody

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody PostTasksJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /collections)
	GetCollections(w http.ResponseWriter, r *http.Request)

	// (GET /collections/{id})
	GetCollectionsId(w http.ResponseWriter, r *http.Request, id CollectionId)

	// (GET /files/{id})
	GetFilesId(w http.ResponseWriter, r *http.Request, id FileIdPathParam)

	// (GET /files/{id}/original/{filename})
	GetFilesIdOriginalFilename(w http.ResponseWriter, r *http.Request, id FileIdPathParam, filename FilenamePathParam)

	// (GET /files/{id}/variants/{size}/{filename})
	GetFilesIdVariantsSizeFilename(w http.ResponseWriter, r *http.Request, id FileIdPathParam, size SizePathParam, filename FilenamePathParam)

	// (GET /scenes)
	GetScenes(w http.ResponseWriter, r *http.Request, params GetScenesParams)

	// (POST /scenes)
	PostScenes(w http.ResponseWriter, r *http.Request)

	// (GET /scenes/{id})
	GetScenesId(w http.ResponseWriter, r *http.Request, id SceneId)

	// (GET /scenes/{scene_id}/dates)
	GetScenesSceneIdDates(w http.ResponseWriter, r *http.Request, sceneId SceneId, params GetScenesSceneIdDatesParams)

	// (GET /scenes/{scene_id}/regions)
	GetScenesSceneIdRegions(w http.ResponseWriter, r *http.Request, sceneId SceneId, params GetScenesSceneIdRegionsParams)

	// (GET /scenes/{scene_id}/regions/{id})
	GetScenesSceneIdRegionsId(w http.ResponseWriter, r *http.Request, sceneId SceneId, id RegionId)

	// (GET /scenes/{scene_id}/tiles)
	GetScenesSceneIdTiles(w http.ResponseWriter, r *http.Request, sceneId SceneId, params GetScenesSceneIdTilesParams)

	// (GET /tasks)
	GetTasks(w http.ResponseWriter, r *http.Request, params GetTasksParams)

	// (POST /tasks)
	PostTasks(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetCollections operation middleware
func (siw *ServerInterfaceWrapper) GetCollections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCollections(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetCollectionsId operation middleware
func (siw *ServerInterfaceWrapper) GetCollectionsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id CollectionId

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCollectionsId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetFilesId operation middleware
func (siw *ServerInterfaceWrapper) GetFilesId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id FileIdPathParam

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFilesId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetFilesIdOriginalFilename operation middleware
func (siw *ServerInterfaceWrapper) GetFilesIdOriginalFilename(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id FileIdPathParam

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "filename" -------------
	var filename FilenamePathParam

	err = runtime.BindStyledParameter("simple", false, "filename", chi.URLParam(r, "filename"), &filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter filename: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFilesIdOriginalFilename(w, r, id, filename)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetFilesIdVariantsSizeFilename operation middleware
func (siw *ServerInterfaceWrapper) GetFilesIdVariantsSizeFilename(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id FileIdPathParam

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "size" -------------
	var size SizePathParam

	err = runtime.BindStyledParameter("simple", false, "size", chi.URLParam(r, "size"), &size)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter size: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "filename" -------------
	var filename FilenamePathParam

	err = runtime.BindStyledParameter("simple", false, "filename", chi.URLParam(r, "filename"), &filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter filename: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFilesIdVariantsSizeFilename(w, r, id, size, filename)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetScenes operation middleware
func (siw *ServerInterfaceWrapper) GetScenes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetScenesParams

	// ------------- Required query parameter "collection_id" -------------
	if paramValue := r.URL.Query().Get("collection_id"); paramValue != "" {

	} else {
		http.Error(w, "Query argument collection_id is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "collection_id", r.URL.Query(), &params.CollectionId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter collection_id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "scene_width" -------------
	if paramValue := r.URL.Query().Get("scene_width"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "scene_width", r.URL.Query(), &params.SceneWidth)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter scene_width: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "image_height" -------------
	if paramValue := r.URL.Query().Get("image_height"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "image_height", r.URL.Query(), &params.ImageHeight)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter image_height: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "layout" -------------
	if paramValue := r.URL.Query().Get("layout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "layout", r.URL.Query(), &params.Layout)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter layout: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetScenes(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostScenes operation middleware
func (siw *ServerInterfaceWrapper) PostScenes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostScenes(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetScenesId operation middleware
func (siw *ServerInterfaceWrapper) GetScenesId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id SceneId

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetScenesId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetScenesSceneIdDates operation middleware
func (siw *ServerInterfaceWrapper) GetScenesSceneIdDates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "scene_id" -------------
	var sceneId SceneId

	err = runtime.BindStyledParameter("simple", false, "scene_id", chi.URLParam(r, "scene_id"), &sceneId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter scene_id: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetScenesSceneIdDatesParams

	// ------------- Required query parameter "height" -------------
	if paramValue := r.URL.Query().Get("height"); paramValue != "" {

	} else {
		http.Error(w, "Query argument height is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "height", r.URL.Query(), &params.Height)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter height: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetScenesSceneIdDates(w, r, sceneId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetScenesSceneIdRegions operation middleware
func (siw *ServerInterfaceWrapper) GetScenesSceneIdRegions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "scene_id" -------------
	var sceneId SceneId

	err = runtime.BindStyledParameter("simple", false, "scene_id", chi.URLParam(r, "scene_id"), &sceneId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter scene_id: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetScenesSceneIdRegionsParams

	// ------------- Required query parameter "x" -------------
	if paramValue := r.URL.Query().Get("x"); paramValue != "" {

	} else {
		http.Error(w, "Query argument x is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "x", r.URL.Query(), &params.X)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter x: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "y" -------------
	if paramValue := r.URL.Query().Get("y"); paramValue != "" {

	} else {
		http.Error(w, "Query argument y is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "y", r.URL.Query(), &params.Y)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter y: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "w" -------------
	if paramValue := r.URL.Query().Get("w"); paramValue != "" {

	} else {
		http.Error(w, "Query argument w is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "w", r.URL.Query(), &params.W)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter w: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "h" -------------
	if paramValue := r.URL.Query().Get("h"); paramValue != "" {

	} else {
		http.Error(w, "Query argument h is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "h", r.URL.Query(), &params.H)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter h: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "limit" -------------
	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter limit: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetScenesSceneIdRegions(w, r, sceneId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetScenesSceneIdRegionsId operation middleware
func (siw *ServerInterfaceWrapper) GetScenesSceneIdRegionsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "scene_id" -------------
	var sceneId SceneId

	err = runtime.BindStyledParameter("simple", false, "scene_id", chi.URLParam(r, "scene_id"), &sceneId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter scene_id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "id" -------------
	var id RegionId

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetScenesSceneIdRegionsId(w, r, sceneId, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetScenesSceneIdTiles operation middleware
func (siw *ServerInterfaceWrapper) GetScenesSceneIdTiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "scene_id" -------------
	var sceneId SceneId

	err = runtime.BindStyledParameter("simple", false, "scene_id", chi.URLParam(r, "scene_id"), &sceneId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter scene_id: %s", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetScenesSceneIdTilesParams

	// ------------- Required query parameter "tile_size" -------------
	if paramValue := r.URL.Query().Get("tile_size"); paramValue != "" {

	} else {
		http.Error(w, "Query argument tile_size is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "tile_size", r.URL.Query(), &params.TileSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter tile_size: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "zoom" -------------
	if paramValue := r.URL.Query().Get("zoom"); paramValue != "" {

	} else {
		http.Error(w, "Query argument zoom is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "zoom", r.URL.Query(), &params.Zoom)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter zoom: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "x" -------------
	if paramValue := r.URL.Query().Get("x"); paramValue != "" {

	} else {
		http.Error(w, "Query argument x is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "x", r.URL.Query(), &params.X)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter x: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "y" -------------
	if paramValue := r.URL.Query().Get("y"); paramValue != "" {

	} else {
		http.Error(w, "Query argument y is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "y", r.URL.Query(), &params.Y)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter y: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "debug_overdraw" -------------
	if paramValue := r.URL.Query().Get("debug_overdraw"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "debug_overdraw", r.URL.Query(), &params.DebugOverdraw)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter debug_overdraw: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "debug_thumbnails" -------------
	if paramValue := r.URL.Query().Get("debug_thumbnails"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "debug_thumbnails", r.URL.Query(), &params.DebugThumbnails)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter debug_thumbnails: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetScenesSceneIdTiles(w, r, sceneId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetTasks operation middleware
func (siw *ServerInterfaceWrapper) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTasksParams

	// ------------- Optional query parameter "type" -------------
	if paramValue := r.URL.Query().Get("type"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "type", r.URL.Query(), &params.Type)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter type: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "collection_id" -------------
	if paramValue := r.URL.Query().Get("collection_id"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "collection_id", r.URL.Query(), &params.CollectionId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter collection_id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTasks(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostTasks operation middleware
func (siw *ServerInterfaceWrapper) PostTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostTasks(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/collections", wrapper.GetCollections)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/collections/{id}", wrapper.GetCollectionsId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/files/{id}", wrapper.GetFilesId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/files/{id}/original/{filename}", wrapper.GetFilesIdOriginalFilename)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/files/{id}/variants/{size}/{filename}", wrapper.GetFilesIdVariantsSizeFilename)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/scenes", wrapper.GetScenes)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/scenes", wrapper.PostScenes)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/scenes/{id}", wrapper.GetScenesId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/scenes/{scene_id}/dates", wrapper.GetScenesSceneIdDates)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/scenes/{scene_id}/regions", wrapper.GetScenesSceneIdRegions)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/scenes/{scene_id}/regions/{id}", wrapper.GetScenesSceneIdRegionsId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/scenes/{scene_id}/tiles", wrapper.GetScenesSceneIdTiles)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/tasks", wrapper.GetTasks)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/tasks", wrapper.PostTasks)
	})

	return r
}

