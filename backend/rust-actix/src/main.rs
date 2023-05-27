extern crate diesel;

use actix_cors::Cors;
use actix_web::{middleware, web, App, HttpServer};
use clap::Parser;
use diesel::prelude::*;
use diesel::r2d2::{self, ConnectionManager};

mod args_parse;
mod controllers;
mod helpers;
mod models;

use args_parse::Args;

pub type DbPool = r2d2::Pool<ConnectionManager<PgConnection>>;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Loading .env into environment variable.
    dotenv::dotenv().ok();

    let args = Args::parse();

    // set up database connection pool
    let database_url = std::env::var("DATABASE_URL").expect("DATABASE_URL");
    let manager = ConnectionManager::<PgConnection>::new(database_url);
    let pool: DbPool = r2d2::Pool::builder()
        .build(manager)
        .expect("Failed to create pool.");

    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));

    println!("Starting server on port {}", args.port);

    HttpServer::new(move || {
        let cors = Cors::permissive();

        App::new()
            .app_data(web::Data::new(pool.clone()))
            .wrap(cors)
            .wrap(middleware::Logger::default())
            .service(
                web::scope("/v1")
                    .service(controllers::healthcheck::index)
                    .service(
                        web::scope("/todos")
                            .service(controllers::todos::index)
                            .service(controllers::todos::show)
                            .service(controllers::todos::create)
                            .service(controllers::todos::update)
                            .service(controllers::todos::delete),
                    ),
            )
    })
    .bind(("0.0.0.0", args.port))?
    .run()
    .await
}
