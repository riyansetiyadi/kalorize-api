CREATE TABLE histories (
  id_history varchar(191) NOT NULL,
  id_user char(36) DEFAULT NULL,
  id_breakfast char(36) DEFAULT NULL,
  id_lunch char(36) DEFAULT NULL,
  id_dinner char(36) DEFAULT NULL,
  total_protein int(11) DEFAULT NULL,
  total_kalori int(11) DEFAULT NULL,
  tanggal_dibuat timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
