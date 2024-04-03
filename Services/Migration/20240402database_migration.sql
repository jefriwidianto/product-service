CREATE DATABASE IF NOT EXISTS diksha_serv_product;

USE diksha_serv_product;

CREATE TABLE IF NOT EXISTS t_product(
    id VARCHAR(36) UNIQUE PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    price FLOAT,
    stock INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME
);

INSERT INTO t_product (id, title, description, price, stock) VALUES ('1a8c926a-35d3-451f-99c2-964cf0666d65', 'BAJU ANAK - kemeja model flanel', '- untuk anak umur 3 tahun - pilihan warna hitam, merah, biru', '350000', '12') ON DUPLICATE KEY UPDATE id = '1a8c926a-35d3-451f-99c2-964cf0666d65';
INSERT INTO t_product (id, title, description, price, stock) VALUES ('937470d6-914e-4326-b904-48c8134219b8', 'Kemeja lengan pendek', '- tersedia untuk ukuran L, dan XL - pilihan warna abu-abu, merah, biru', '450000', '134') ON DUPLICATE KEY UPDATE id = '937470d6-914e-4326-b904-48c8134219b8';
INSERT INTO t_product (id, title, description, price, stock) VALUES ('545831f2-46ad-4845-b794-afff3db72305', 'Oli motor', '- Oli synthetic untuk sepeda motor', '50000', '10') ON DUPLICATE KEY UPDATE id = '545831f2-46ad-4845-b794-afff3db72305'