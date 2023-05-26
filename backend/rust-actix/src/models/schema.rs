// @generated automatically by Diesel CLI.

diesel::table! {
    schema_migrations (version) {
        version -> Int8,
        dirty -> Bool,
    }
}

diesel::table! {
    todos (id) {
        id -> Int4,
        title -> Varchar,
        completed -> Bool,
        order -> Int4,
        url -> Varchar,
        created_at -> Timestamp,
        updated_at -> Timestamp,
    }
}

diesel::allow_tables_to_appear_in_same_query!(
    schema_migrations,
    todos,
);
