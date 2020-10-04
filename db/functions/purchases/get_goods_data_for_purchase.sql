SET search_path TO budget;

CREATE OR REPLACE FUNCTION get_goods_data_for_purchase(purchase_id INTEGER)
    RETURNS TABLE(goods_details_id INTEGER,
                  amount NUMERIC,
                  price MONEY,
                  goods_id INTEGER,
                  name VARCHAR(3000),
                  category_id INTEGER,
                  comment VARCHAR(3000),
                  bar_code VARCHAR(100)) AS
$$
SELECT goods_details.id AS goods_details_id,
       amount,
       price,
       goods.id AS goods_id,
       name,
       category_id,
       comment,bar_code
FROM goods_details
         JOIN goods ON goods.id = goods_details.goods_id
WHERE goods_details.purchase_id = purchase_id
  AND goods_details.is_removed = FALSE
  AND goods.is_removed = FALSE
$$
    LANGUAGE sql;