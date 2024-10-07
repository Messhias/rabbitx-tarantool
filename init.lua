box.cfg{}

if not box.space.positions then
    box.schema.space.create('positions', {
        format = {
            {name = 'userID', type = 'unsigned'},
            {name = 'market', type = 'string'},
            {name = 'size', type = 'number'},
            {name = 'entryPrice', type = 'number'},
            {name = 'side', type = 'string'}
        }
    })
    box.space.positions:create_index('primary', {
        parts = {'userID', 'market'}
    })
end

