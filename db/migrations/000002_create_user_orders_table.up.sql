CREATE TABLE user_orders
(
    order_number text NOT NULL,
    user_id bigint NOT NULL,
    placed_at timestamp with time zone NOT NULL,
    status character varying(256) NOT NULL,
    PRIMARY KEY (order_number)
);

ALTER TABLE user_orders
    ADD CONSTRAINT fk_user_orders_user_id FOREIGN KEY (user_id)
        REFERENCES users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE;
