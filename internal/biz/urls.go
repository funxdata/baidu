package biz

const (
	//人体检测
	URLBodyTracking = "/rest/2.0/image-classify/v1/body_tracking"
	// h5语音验证码
	URLFacesessioncode = "/rest/2.0/face/v1/faceliveness/sessioncode"
	// h5活体视频分析
	URLFaceVerify = "/rest/2.0/face/v1/faceliveness/verify"
	// 人脸检测
	URLFaceDetect = "/rest/2.0/face/v3/detect"
	// 人脸对比
	URLFaceMatch = "/rest/2.0/face/v3/match"
	// 人脸搜索
	URLFaceSearch = "/rest/2.0/face/v3/search"
	// 人脸库管理-人脸注册
	URLFaceUserAdd = "/rest/2.0/face/v3/faceset/user/add"
	// 人脸库管理-人脸更新
	URLFaceUserUpdate = "/rest/2.0/face/v3/faceset/user/update"
	// 人脸库管理-删除用户
	URLFaceUserDelete = "/rest/2.0/face/v3/faceset/user/delete"
	// 人脸库管理-用户信息查询
	URLFaceUserGet = "/rest/2.0/face/v3/faceset/user/get"
	// 人脸库管理-获取组列表
	URLFaceGroups = "/rest/2.0/face/v3/faceset/group/getlist"
	// 人脸库管理-获取用户列表
	URLFaceGroupUsers = "/rest/2.0/face/v3/faceset/group/getusers"
	// 人脸库管理-复制用户
	URLFaceCopy = "/rest/2.0/face/v3/faceset/user/copy"
	// 人脸库管理-获取用户人脸列表
	URLFaces = "/rest/2.0/face/v3/faceset/face/getlist"
	// 人脸库管理-创建用户组
	URLFaceGroupAdd = "/rest/2.0/face/v3/faceset/group/add"
	// 人脸库管理-删除用户组
	URLFaceGroupDelete = "/rest/2.0/face/v3/faceset/group/delete"
	// 人脸库管理-删除人脸
	URLFaceDelete = "/rest/2.0/face/v3/faceset/face/delete"
	// 在线活体检测
	URLFaceFaceverify = "/rest/2.0/face/v3/faceverify"
)
