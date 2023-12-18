
-- +migrate Up
CREATE TABLE `users` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT 'ユーザ名',
  `email` varchar(255) NOT NULL COMMENT 'メールアドレス',
  `firebase_uid` varchar(255) NOT NULL COMMENT 'Firebaseが持っているID',
  `image` varchar(255) NOT NULL COMMENT '表示用写真URL',
  `created_at` timestamp,
  `updated_at` timestamp
);

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
);

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
);

CREATE TABLE `bookmarks` (
  `event_id` INT COMMENT 'イベントID',
  `user_id` INT COMMENT 'ユーザID',
  PRIMARY KEY (`event_id`, `user_id`)
);


-- +migrate Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS bookmarked_events;
DROP TABLE IF EXISTS bookmarks;