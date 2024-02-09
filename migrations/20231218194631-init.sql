
-- +migrate Up
CREATE TABLE `users` (
  `id` VARCHAR(255) NOT NULL COMMENT 'ユーザ識別子',
  `name` VARCHAR(255) COMMENT 'ユーザ名',
  `email` VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
  `emailVerified` DATETIME(6) COMMENT 'メールアドレスの認証日時',
  `image` VARCHAR(255) COMMENT '表示用写真URL',
  `created_at` timestamp,
  `updated_at` timestamp,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_email` (`email`)
) COMMENT 'Auth.jsでユーザの識別子を管理するテーブル';

CREATE TABLE `accounts` (
  `userId` varchar(255) NOT NULL COMMENT 'userテーブルのid、ユーザーの識別子',
  `type` varchar(255) NOT NULL COMMENT 'oauth等の値が入る',
  `provider` varchar(255) NOT NULL COMMENT 'google, line等の認証Providerの名称が入る',
  `providerAccountId` varchar(255) NOT NULL COMMENT 'google, line等の認証Providerのユーザー識別子が入る',
  `refresh_token` text COMMENT '認証Providerから取得したリフレッシュトークンが記録される',
  `access_token` text COMMENT '認証Providerから取得したアクセストークンが記録される',
  `expires_at` int DEFAULT NULL COMMENT '認証Providerから取得したアクセストークンの有効期限',
  `token_type` varchar(255) DEFAULT NULL COMMENT 'Bearer等のトークン種別が記録される',
  `scope` varchar(255) DEFAULT NULL COMMENT 'リクエスト時に認証Providerに要求したscopeが記録される',
  `id_token` text COMMENT '認証Providerから取得したIDトークンの有効期限',
  `session_state` varchar(255) DEFAULT NULL,
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`provider`,`providerAccountId`),
  UNIQUE KEY `uk_account_userid` (`userId`),
  KEY `idx_account_01` (`userId`)
) COMMENT 'Auth.jsでユーザの認証情報を管理するテーブル';

CREATE TABLE `sessions` (
  `sessionToken` varchar(255) NOT NULL COMMENT 'セッションの識別子が格納される',
  `userId` varchar(255) NOT NULL COMMENT 'userテーブルのid、ユーザーの識別子',
  `expires` timestamp NOT NULL COMMENT 'セッションの有効期限が格納される',
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`sessionToken`),
  UNIQUE KEY `uk_session_userid` (`userId`),
  KEY `idx_session_01` (`userId`)
) COMMENT 'Auth.jsでユーザーのセッションを管理するテーブル session.strategyがdatabaseの時だけ利用される';

CREATE TABLE `verificationtokens` (
  `identifier` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `expires` timestamp NOT NULL,
  PRIMARY KEY (`identifier`,`token`)
) COMMENT 'Auth.jsでメールアドレスを使ったパスワードレスログインを利用する時に利用される';

CREATE TABLE `events` (
  `event_id` INT PRIMARY KEY NOT NULL COMMENT 'イベントID',
  `title` VARCHAR(255) NOT NULL COMMENT 'タイトル',
  `catch` VARCHAR(255) COMMENT 'キャッチ文章',
  `description` TEXT COMMENT '概要',
  `event_url` VARCHAR(255) COMMENT 'イベントURL',
  `started_at` DATETIME COMMENT '開始日時',
  `ended_at` DATETIME COMMENT '終了日時',
  `limit` INT COMMENT '定員',
  `hash_tag` VARCHAR(100) COMMENT 'ハッシュタグ',
  `event_type` VARCHAR(50) COMMENT 'イベントタイプ',
  `accepted` INT COMMENT '参加人数',
  `waiting` INT COMMENT '補欠人数',
  `updated_at` DATETIME COMMENT '更新日時',
  `owner_id` INT COMMENT '管理者ID',
  `owner_nickname` VARCHAR(255) COMMENT '管理者ニックネーム',
  `owner_display_name` VARCHAR(255) COMMENT '管理者名',
  `place` VARCHAR(255) COMMENT '開催場所',
  `address` VARCHAR(255) COMMENT '開催住所',
  `lat` VARCHAR(255) NOT NULL COMMENT '緯度',
  `lon` VARCHAR(255) NOT NULL COMMENT '経度'
) COMMENT 'イベント情報を管理するテーブル';

CREATE TABLE `bookmarked_events` (
  `event_id` INT PRIMARY KEY NOT NULL COMMENT 'イベントID',
  `title` VARCHAR(255) NOT NULL COMMENT 'タイトル',
  `catch` VARCHAR(255) COMMENT 'キャッチ文章',
  `description` TEXT COMMENT '概要',
  `event_url` VARCHAR(255) COMMENT 'イベントURL',
  `started_at` DATETIME COMMENT '開始日時',
  `ended_at` DATETIME COMMENT '終了日時',
  `limit` INT COMMENT '定員',
  `hash_tag` VARCHAR(100) COMMENT 'ハッシュタグ',
  `event_type` VARCHAR(50) COMMENT 'イベントタイプ',
  `accepted` INT COMMENT '参加人数',
  `waiting` INT COMMENT '補欠人数',
  `updated_at` DATETIME COMMENT '更新日時',
  `owner_id` INT COMMENT '管理者ID',
  `owner_nickname` VARCHAR(255) COMMENT '管理者ニックネーム',
  `owner_display_name` VARCHAR(255) COMMENT '管理者名',
  `place` VARCHAR(255) COMMENT '開催場所',
  `address` VARCHAR(255) COMMENT '開催住所',
  `lat` VARCHAR(255) NOT NULL COMMENT '緯度',
  `lon` VARCHAR(255) NOT NULL COMMENT '経度'
) COMMENT 'ブックマークされたイベントを管理するテーブル';

CREATE TABLE `bookmarks` (
  `event_id` INT COMMENT 'イベントID',
  `user_id` INT COMMENT 'ユーザID',
  PRIMARY KEY (`event_id`, `user_id`)
) COMMENT 'ユーザとユーザがブックマークしたイベントを紐づけるテーブル';


-- +migrate Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS verificationtokens;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS bookmarked_events;
DROP TABLE IF EXISTS bookmarks;