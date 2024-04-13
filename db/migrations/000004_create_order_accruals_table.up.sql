CREATE TABLE order_accruals
(
    order_number text NOT NULL,
    user_id bigint NOT NULL,
    status character varying(256) NOT NULL,
    amount numeric,
    CONSTRAINT order_accruals_pkey PRIMARY KEY (order_number)
);

ALTER TABLE order_accruals
    ADD CONSTRAINT fk_order_accruals_user_id FOREIGN KEY (user_id)
        REFERENCES users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE;

ALTER TABLE order_accruals
    ADD CONSTRAINT fk_order_accruals_order_number FOREIGN KEY (order_number)
        REFERENCES user_orders (order_number) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE;
