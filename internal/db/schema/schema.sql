CREATE TABLE brands
(
    id         BIGINT AUTO_INCREMENT COMMENT 'pk'
PRIMARY KEY,
    name       VARCHAR(100)                       NOT NULL COMMENT '브랜드명',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '생성일',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '갱신일',
    deleted_at DATETIME NULL COMMENT '삭제일'
) COMMENT '브랜드';

CREATE TABLE products
(
    id          BIGINT AUTO_INCREMENT COMMENT 'pk'
PRIMARY KEY,
    brand_id    BIGINT                                NOT NULL COMMENT '브랜드 id',
    name        VARCHAR(255)                          NOT NULL COMMENT '상품명',
    description TEXT NULL COMMENT '설명',
    base_price  INT                                   NOT NULL COMMENT '정가',
    sale_status VARCHAR(20) DEFAULT 'ON_SALE'         NOT NULL COMMENT '세일 상태( ON_SALE: 세일중 )',
    created_at  DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '생성일',
    updated_at  DATETIME    DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '갱신일',
    deleted_at DATETIME NULL COMMENT '삭제일',
    CONSTRAINT fk_brand_id
        FOREIGN KEY (brand_id) REFERENCES brands (id)
) COMMENT '상품';

CREATE INDEX idx_brand_id
    ON products (brand_id, base_price) COMMENT '브랜드별 가격 정렬';