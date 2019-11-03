INSERT INTO games
    (
    game_complete,
    game_state,
    pending_moves
    )
VALUES
    (
        FALSE,
        '{"id": 3,"turn_number": 0,"land": [{"id": 0,"owner_id": 0,"crop": 100,"infected_crop": 0},{"id": 1,"owner_id": -1,"crop": 100,"infected_crop": 0},{"id": 2,"owner_id": -1,"crop": 100,"infected_crop": 0},{"id": 3,"owner_id": -1,"crop": 100,"infected_crop": 0},{"id": 4,"owner_id": 1,"crop": 100,"infected_crop": 0},{"id": 5,"owner_id": -1,"crop": 100,"infected_crop": 0},{"id": 6,"owner_id": -1,"crop": 100,"infected_crop": 0},{"id": 7,"owner_id": -1,"crop": 100,"infected_crop": 0},{"id": 8,"owner_id": 2,"crop": 10,"infected_crop": 90},{"id": 9,"owner_id": -1,"crop": 100,"infected_crop": 0},{"id": 10,"owner_id": -1,"crop": 100,"infected_crop": 0},{"id": 11,"owner_id": -1,"crop": 100,"infected_crop": 0},{"id": 12,"owner_id": 3,"crop": 100,"infected_crop": 0}],"players": [{"id": 0,"cash": 10000,"ready": false,"land_color": "#D24858"},{"id": 1,"cash": 10000,"ready": false,"land_color": "#EA8676"},{"id": 2,"cash": 10000,"ready": false,"land_color": "#EAB05E"},{"id": 3,"cash": 10000,"ready": false,"land_color": "#FDEECD"}],"traders": [{"id": "ab34","owner_id": 0,"safe_crop": 50,"infected_crop": 10,"location_id": 0,"destination_id": 10},{"id": "56ty","owner_id": -1,"safe_crop": 40,"infected_crop": 0,"location_id": 2,"destination_id": 12}]}',
        '{"buy_crop": [{"id": "abc","player_id": 0,"land_id": 0,"trader_id": "ab34" }],"create_trader": [{"id": "def","player_id": 1,"safe_crop": 50,"infected_crop": 50,"land_id": 4}],"route_trader": [{"id": "123","player_id": 0,"trader_id": "ab34","destination_id": 6}],"infect_crop": [{"id": "hij","player_id": 2,"amount": 10,"land_id": 0}],"buy_land": [{"id": "klm","player_id": 0,"land_id": 1,"bid_price":400}]}'
);