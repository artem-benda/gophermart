CREATE TABLE order_withdrawals
(
    order_number character varying(256),
    user_id bigint NOT NULL,
    amount numeric NOT NULL,
    created_at timestamp with time zone NOT NULL,
    processed_at timestamp with time zone,
    CONSTRAINT order_withdrawals_pkey PRIMARY KEY (order_number)
);

ALTER TABLE order_withdrawals
    ADD CONSTRAINT fk_order_withdrawals_user_id FOREIGN KEY (user_id)
        REFERENCES users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE;

/*
ALTER TABLE order_withdrawals
    ADD CONSTRAINT fk_order_withdrawals_order_number FOREIGN KEY (order_number)
        REFERENCES user_orders (order_number) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE;
*/
