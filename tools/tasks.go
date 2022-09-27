package tools

import (
	"encoding/json"
)

func AuthUser() (User, error) {
	res, err := SendRequest(RequestInfo{ReqType: "GET", Path: "/user"})
	if err != nil {
		return User{}, err
	}
	var user User
	json.Unmarshal(res.ResponseInfo, &user)

	return user, nil
}

func GetRepos() ([]RepoObject, error) {

	res, err := SendRequest(RequestInfo{ReqType: "GET", Path: "/user/repos"})
	if err != nil {
		return nil, err
	}
	var reposArr []RepoObject

	json.Unmarshal(res.ResponseInfo, &reposArr)

	return reposArr, nil

}

func CreateRepo(repoInfo NewRepo) (NewRepo, error) {
	body := &repoInfo
	// serialize object

	out, err := json.Marshal(body)
	if err != nil {
		println(err)
	}
	res, err := SendRequest(RequestInfo{ReqType: "POST", Path: "/user/repos", Body: string([]byte(out))})
	if err != nil {
		return NewRepo{}, err
	}
	var newRepo NewRepo

	json.Unmarshal(res.ResponseInfo, &newRepo)

	return newRepo, err
}
