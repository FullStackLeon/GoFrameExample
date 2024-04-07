CREATE TABLE users (
   user_id INT AUTO_INCREMENT PRIMARY KEY, -- 用户ID，自增主键
   username VARCHAR(255) NOT NULL, -- 用户名，不能为空
   email VARCHAR(255) UNIQUE, -- 电子邮箱地址，唯一索引
   password VARCHAR(255) NOT NULL, -- 加密后的密码，长度一般为哈希后的结果，不能为空
   account_type ENUM('regular', 'auth', 'admin') DEFAULT 'regular', -- 账号类型：普通用户、auth用户、管理员用户，默认为普通用户
   auth_provider VARCHAR(50), -- auth提供商名称，如Google、Twitter等
   auth_user_id VARCHAR(255), -- 用户在auth提供商系统中的唯一标识符
   auth_access_token VARCHAR(255), -- auth访问令牌
   auth_refresh_token VARCHAR(255), -- auth刷新令牌
   auth_user_info JSON, -- auth提供商返回的用户信息，JSON格式
   balance DECIMAL(16, 2) DEFAULT 0.00, -- 用户余额，DECIMAL类型，精确到小数点后两位，默认为0.00
   source VARCHAR(255), -- 用户来源信息，如Twitter等平台
   status ENUM('active', 'inactive', 'frozen', 'suspended', 'deleted') DEFAULT 'active', -- 用户状态，active为激活状态，inactive为未激活状态，frozen为已冻结状态，suspended为已暂停状态，deleted为已删除状态，默认为active
   google_auth_key VARCHAR(255), -- 用户Google Authenticator的密钥
   two_factor_auth ENUM('enabled', 'disabled') DEFAULT 'disabled', -- 用户两步验证状态，enabled为已启用，disabled为已禁用，默认为disabled
   kyc_status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending', -- KYC（Know Your Customer）状态，pending为待审核，approved为已审核通过，rejected为审核未通过，默认为pending
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 记录创建时间，默认为当前时间
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 记录更新时间，默认为当前时间，自动更新
   deleted_at TIMESTAMP -- 记录删除时间，用于软删除，为空表示未删除
)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;