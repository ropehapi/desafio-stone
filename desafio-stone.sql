-- MySQL Script generated by MySQL Workbench
-- Wed Jun  5 13:26:49 2024
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema desafio_stone
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `desafio_stone` ;

-- -----------------------------------------------------
-- Schema desafio_stone
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `desafio_stone` DEFAULT CHARACTER SET utf8 ;
USE `desafio_stone` ;

-- -----------------------------------------------------
-- Table `desafio_stone`.`person`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `desafio_stone`.`person` ;

CREATE TABLE IF NOT EXISTS `desafio_stone`.`person` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `desafio_stone`.`relationship`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `desafio_stone`.`relationship` ;

CREATE TABLE IF NOT EXISTS `desafio_stone`.`relationship` (
  `children_id` INT NOT NULL,
  `parent_1_id` INT NOT NULL,
  `parent_2_id` INT NOT NULL,
  PRIMARY KEY (`children_id`, `parent_1_id`, `parent_2_id`),
  INDEX `fk_relationship_person_idx` (`children_id` ASC),
  INDEX `fk_relationship_person1_idx` (`parent_1_id` ASC),
  INDEX `fk_relationship_person2_idx` (`parent_2_id` ASC),
  CONSTRAINT `fk_relationship_person`
    FOREIGN KEY (`children_id`)
    REFERENCES `desafio_stone`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_relationship_person1`
    FOREIGN KEY (`parent_1_id`)
    REFERENCES `desafio_stone`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_relationship_person2`
    FOREIGN KEY (`parent_2_id`)
    REFERENCES `desafio_stone`.`person` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
