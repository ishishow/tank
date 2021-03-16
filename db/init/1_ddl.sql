-- MySQL Script generated by MySQL Workbench
-- Fri Feb 14 23:09:20 2020
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema dojo_api
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema dojo_api
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `dojo_api` DEFAULT CHARACTER SET utf8mb4 ;
USE `dojo_api` ;

SET CHARSET utf8mb4;

-- -----------------------------------------------------
-- Table `dojo_api`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dojo_api`.`user` (
  `id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
  `auth_token` VARCHAR(128) NOT NULL COMMENT '認証トークン',
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  `high_score` INT UNSIGNED NOT NULL COMMENT 'ハイスコア',
  `coin` INT UNSIGNED NOT NULL COMMENT '所持コイン',
  PRIMARY KEY (`id`),
  INDEX `idx_auth_token` (`auth_token` ASC))
ENGINE = InnoDB
COMMENT = 'ユーザ';


-- -----------------------------------------------------
-- Table `dojo_api`.`item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dojo_api`.`item` (
  `id` VARCHAR(128) NOT NULL COMMENT 'アイテムID',
  `name` VARCHAR(64) NOT NULL COMMENT 'アイテム名',
  `rarity` INT NOT NULL COMMENT 'レアリティ',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'アイテム';

-- -----------------------------------------------------
-- Table `dojo_api`.`weapon`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dojo_api`.`weapon` (
  `id` VARCHAR(128) NOT NULL COMMENT 'ウェポンID',
  `item_id` VARCHAR(128) NOT NULL COMMENT 'アイテムID',
  `ballet` INT NOT NULL COMMENT '弾丸',
  `attack` INT NOT NULL COMMENT '攻撃力',
  `reload` FLOAT NOT NULL COMMENT 'リロード',
  `speed` FLOAT NOT NULL COMMENT '弾丸速度',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = '武器';

-- -----------------------------------------------------
-- Table `dojo_api`.`skin`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dojo_api`.`skin` (
  `id` VARCHAR(128) NOT NULL COMMENT 'スキンID',
  `item_id` VARCHAR(128) NOT NULL COMMENT 'アイテムID',
  `hit_point` INT NOT NULL COMMENT 'ヒットポイント',
  `speed` FLOAT NOT NULL COMMENT '移動速度',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'スキン';

-- -----------------------------------------------------
-- Table `dojo_api`.`user_status`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dojo_api`.`user_status` (
  `user_id` VARCHAR(128) NOT NULL COMMENT 'アイテムID',
  `weapon_id` VARCHAR(128) NOT NULL COMMENT 'アイテムID',
  `skin_id` VARCHAR(128) NOT NULL COMMENT 'アイテムID')
ENGINE = InnoDB
COMMENT = 'ユーザーステータス';

-- -----------------------------------------------------
-- Table `dojo_api`.`user_item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dojo_api`.`user_item` (
  `user_id` VARCHAR(128) NOT NULL COMMENT 'アイテムID',
  `item_id` VARCHAR(128) NOT NULL COMMENT 'アイテムID')
ENGINE = InnoDB
COMMENT = 'ユーザーステータス';

-- -----------------------------------------------------
-- Table `dojo_api`.`gacha_probability`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `dojo_api`.`gacha_probability` (
  `item_id` VARCHAR(128) NOT NULL COMMENT 'アイテムID',
  `ratio` INT UNSIGNED NOT NULL COMMENT '排出重み')
ENGINE = InnoDB
COMMENT = 'ガチャ排出情報';


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
