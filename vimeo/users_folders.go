package vimeo

import (
	"fmt"
	"time"
)

type dataListFolder struct {
	Data []*Folder `json:"data"`
	pagination
}


// Folder represents a folder.
type Folder struct {
	URI          string    `json:"uri,omitempty"`
	Name         string    `json:"name,omitempty"`
	Link         string    `json:"link,omitempty"`
	CreatedTime  time.Time `json:"created_time,omitempty"`
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	User         *User     `json:"user,omitempty"`
	Privacy      *Privacy  `json:"privacy,omitempty"`
}

// ListFolder method gets all the folders from the specified user's account.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_folders
func (s *UsersService) ListFolder(uid string, opt ...CallOption) ([]*Folder, *Response, error) {
	var u string
	if uid == "" {
		u = "me/projects"
	} else {
		u = fmt.Sprintf("users/%s/projects", uid)
	}

	u, err := addOptions(u, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	folders := &dataListFolder{}

	resp, err := s.client.Do(req, folders)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(folders)

	return folders.Data, resp, err
}

// GetFolder method gets a single folder.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_folder
func (s *UsersService) GetFolder(uid string, ab string, opt ...CallOption) (*Folder, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/projects/%s", ab)
	} else {
		u = fmt.Sprintf("users/%s/projects/%s", uid, ab)
	}

	u, err := addOptions(u, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	folder := &Folder{}

	resp, err := s.client.Do(req, folder)
	if err != nil {
		return nil, resp, err
	}

	return folder, resp, err
}

// FolderListVideo method gets all the videos from the specified folder.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_folder_videos
func (s *UsersService) FolderListVideo(uid string, ab string, opt ...CallOption) ([]*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/projects/%s/videos", ab)
	} else {
		u = fmt.Sprintf("users/%s/projects/%s/videos", uid, ab)
	}
	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// FolderGetVideo method gets a single video from an folder. You can use this method to determine whether the folder contains the specified video.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_folder_video
func (s *UsersService) FolderGetVideo(uid string, ab string, vid int, opt ...CallOption) (*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/projects/%s/videos/%d", ab, vid)
	} else {
		u = fmt.Sprintf("users/%s/projects/%s/videos/%d", uid, ab, vid)
	}
	video, resp, err := getVideo(s.client, u, opt...)

	return video, resp, err
}
