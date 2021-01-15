class CreateUsers < ActiveRecord::Migration[5.1]
  def up
    create_table :users do |t|
      t.string :name, limit: 50, null: false
      t.string :phone, unique: true, limit: 20, null: false
      t.string :password, null: false
      t.string :role, null: false
      t.timestamps

      t.index :phone
    end
  end

  def down
    drop_table :users
  end
end
