DELETE FROM tb_halls WHERE theater_id in (
    1,
    2,
)

DELETE FROM tb_theaters WHERE name IN('大地影院', '北极星影城');
