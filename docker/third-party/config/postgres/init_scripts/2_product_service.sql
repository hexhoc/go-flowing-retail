-- ------------------------------------------------------------
-- Dump of table categories
-- ------------------------------------------------------------

drop table if exists "categories";
create table "categories"
(
    id            serial4,
    level         int2         not null,
    parent_id     int4,
    name          varchar(100) not null,
    rank          int          not null,
    is_deleted    boolean      not null,
    created_at  timestamp    not null,
    updated_at timestamp    not null
);
alter table "categories"
    add constraint "category_pk" primary key ("id");

comment on column categories.id is 'Category ID';
comment on column categories.level is 'Category level (1-first-level classification 2-secondary classification 3-third-level classification)';
comment on column categories.parent_id is 'parent category id';
comment on column categories.name is 'Category name';
comment on column categories.rank is 'Sort value (the bigger the field, the higher the front)';
comment on column categories.is_deleted is 'Delete identification field (false-not deleted true-deleted)';
comment on column categories.created_at is 'Creation time';
comment on column categories.updated_at is 'Modification time';

-- ------------------------------------------------------------
-- Dump of table products
-- ------------------------------------------------------------

drop table if exists "products";
create table "products"
(
    id             serial4,
    name           varchar(300) not null default '',
    intro          varchar(500) not null default '',
    description    varchar(500) not null default '',
    category_id    int4         not null,
    original_price numeric      not null default 0 check ( original_price > 0 ),
    selling_price  numeric      not null default 0 check ( selling_price > 0 ),
    is_sale        boolean      not null default false,
    is_deleted     boolean      not null,
    created_at   timestamp    not null,
    updated_at  timestamp    not null
);

alter table products
    add constraint "product_id" primary key (id);
alter table products
    add constraint "product_fk" foreign key (category_id)
        references categories (id)
        on delete set null;

comment on column products.name is 'Product name';
comment on column products.intro is 'Product introduction';
comment on column products.description is 'Product details';
comment on column products.category_id is 'Associated category id';
comment on column products.original_price is 'Product purchase price';
comment on column products.selling_price is 'Product selling price';
comment on column products.is_sale is 'True if the product is on the shelf, and false if it is removed from the shelf';
comment on column products.is_deleted is 'Delete identification field (false-not deleted true-deleted)';
comment on column products.created_at is 'Creation time';
comment on column products.updated_at is 'Modification time';


-- ------------------------------------------------------------
-- Dump of table images
-- ------------------------------------------------------------

drop table if exists product_images;
create table product_images
(
    id          serial4 not null,
    product_id  int4 not null,
    name        varchar(255) not null,
    image_bytes bytea
);

alter table product_images
    add constraint image_pk primary key (id);

alter table product_images
    add constraint product_fk foreign key (product_id)
        references products (id)
        on delete cascade;

comment on column product_images.id is 'image id';
comment on column product_images.product_id is 'product id';
comment on column product_images.name is 'image name (or full path)';
comment on column product_images.image_bytes is 'image binary';


-- Mock data for categories table
INSERT INTO categories (level, parent_id, name, rank, is_deleted, created_at, updated_at)
VALUES (1, null, 'Electronics', 1, false, '2022-10-01 10:00:00', '2022-10-01 10:00:00'),
       (2, 1, 'Computers', 1, false, '2022-10-02 10:00:00', '2022-10-02 10:00:00'),
       (2, 1, 'Smartphones', 2, false, '2022-10-03 10:00:00', '2022-10-03 10:00:00'),
       (3, 2, 'Laptops', 1, false, '2022-10-04 10:00:00', '2022-10-04 10:00:00'),
       (3, 2, 'Desktops', 2, false, '2022-10-05 10:00:00', '2022-10-05 10:00:00'),
       (3, 3, 'Apple', 1, false, '2022-10-06 10:00:00', '2022-10-06 10:00:00'),
       (3, 3, 'Samsung', 2, false, '2022-10-07 10:00:00', '2022-10-07 10:00:00');

INSERT INTO products (name, intro, description, category_id, original_price, selling_price, is_sale, is_deleted, created_at, updated_at)
VALUES ('MacBook Pro', 'Powerful laptop from Apple', 'The MacBook Pro features a stunning Retina display, advanced processors, superfast graphics, and more.', 4, 1599.00, 1799.00, false, false, '2022-10-01 10:00:00', '2022-10-01 10:00:00'),
       ('Galaxy S21', 'Samsung flagship smartphone', 'The Galaxy S21 features a 6.2-inch AMOLED display, 5G connectivity, and a powerful Exynos 2100 processor.', 7, 799.00, 899.00, true, false, '2022-10-02 10:00:00', '2022-10-02 10:00:00'),
       ('iMac', 'All-in-one desktop from Apple', 'The iMac features a beautiful 24-inch Retina display, powerful Apple M1 chip, and a sleek, modern design.', 5, 1299.00, 1499.00, false, false, '2022-10-03 10:00:00', '2022-10-03 10:00:00'),
       ('Galaxy Book Pro', 'Samsung ultrabook', 'The Galaxy Book Pro is an ultra-thin and light laptop with a stunning AMOLED display, 11th Gen Intel Core processor, and 5G connectivity.', 4, 1199.00, 1399.00, true, false, '2022-10-04 10:00:00', '2022-10-04 10:00:00'),
       ('Mac mini', 'Powerful small desktop from Apple', 'The Mac mini is a compact desktop with the latest Apple M1 chip, up to 16GB of RAM, and plenty of ports for connectivity.', 5, 699.00, 799.00, false, false, '2022-10-05 10:00:00', '2022-10-05 10:00:00'),
       ('Galaxy Tab S7', 'Samsung Android tablet', 'The Galaxy Tab S7 features a 12.4-inch Super AMOLED display, 5G connectivity, and the powerful Snapdragon 865+ processor.', 7, 649.00, 749.00, true, false, '2022-10-06 10:00:00', '2022-10-06 10:00:00'),
       ('MacBook Air', 'Thin and light laptop from Apple', 'The MacBook Air is a fan-favorite, with a 13.3-inch Retina display, Apple M1 chip, and up to 18 hours of battery life.', 4, 999.00, 1199.00, false, false, '2022-10-07 10:00:00', '2022-10-07 10:00:00');
