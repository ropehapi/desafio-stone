-- MySQL Script generated by MySQL Workbench
-- Wed Jun  5 15:34:09 2024
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS, UNIQUE_CHECKS = 0;
SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0;
SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE =
        'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema desafio_stone_test
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `desafio_stone_test` DEFAULT CHARACTER SET utf8;
USE `desafio_stone_test`;

-- -----------------------------------------------------
-- Table `desafio_stone_test`.`person`
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS `desafio_stone_test`.`person`
(
    `id`   VARCHAR(36) NOT NULL,
    `name` VARCHAR(45) NOT NULL,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB;

INSERT INTO `desafio_stone`.`person` (id, name) values (uuid(), 'Pedro Yoshimura'), (uuid(), 'Pietra Yoshimura');

-- -----------------------------------------------------
-- Table `desafio_stone_test`.`relationship`
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS `desafio_stone_test`.`relationship`
(
    `children_id` VARCHAR(36) NOT NULL,
    `parent_id`   VARCHAR(36) NOT NULL,
    PRIMARY KEY (`children_id`, `parent_id`),
    INDEX `fk_relationship_person1_idx` (`parent_id` ASC),
    CONSTRAINT `fk_relationship_person`
        FOREIGN KEY (`children_id`)
            REFERENCES `desafio_stone_test`.`person` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `fk_relationship_person1`
        FOREIGN KEY (`parent_id`)
            REFERENCES `desafio_stone_test`.`person` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;


SET SQL_MODE = @OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS;
