-- Расширение для генерации UUID


-- === Таблица заказов (основная) ===
CREATE TABLE IF NOT EXISTS orders (
                                      order_uid         text PRIMARY KEY,
                                      track_number      text       ,
                                      entry             varchar(100) ,
                                      locale            varchar(5),
                                      internal_signature text,
                                      customer_id       text,
                                      delivery_service  varchar(100),
                                      shard_key         varchar(10),
                                      sm_id             integer,
                                      date_created      timestamptz NOT NULL,
                                      oof_shard         varchar(10)
);




CREATE TABLE IF NOT EXISTS deliveries (
                                          id         uuid PRIMARY KEY,
                                          order_uid  text UNIQUE NOT NULL REFERENCES orders(order_uid) ON DELETE CASCADE,
                                          name       text,
                                          phone      text,
                                          zip        text,
                                          city       text,
                                          address    text,
                                          region     text,
                                          email      text
);

-- === Таблица оплаты (1:1 с заказом) ===
CREATE TABLE IF NOT EXISTS payments (
                                        id             uuid PRIMARY KEY,
                                        order_uid      text UNIQUE NOT NULL REFERENCES orders(order_uid) ON DELETE CASCADE,
                                        transaction    text,
                                        request_id     text,
                                        currency       text,
                                        provider       text,
                                        amount         integer,
                                        payment_dt     timestamptz,
                                        bank           text,
                                        delivery_cost  integer,
                                        goods_total    integer,
                                        custom_fee     integer
);



CREATE TABLE IF NOT EXISTS items (
                                     id           uuid PRIMARY KEY ,
                                     order_uid    text NOT NULL REFERENCES orders(order_uid) ON DELETE CASCADE,
                                     chrt_id      integer NOT NULL,
                                     track_number text,
                                     price        integer,
                                     rid          text,
                                     name         text,
                                     sale         integer,
                                     size         text,
                                     total_price  integer,
                                     nm_id        integer,
                                     brand        text,
                                     status       integer,
                                     UNIQUE(order_uid, chrt_id)
);


