using Microsoft.EntityFrameworkCore;
using Recipes.API.db;
using Recipes.API.extensions;
using Recipes.API.handlers;
using Recipes.API.routes;

var builder = WebApplication.CreateBuilder(args);

// Services
builder.Services.AddApplicationServices();
builder.Services.AddAuthorizationServices(builder.Configuration);

// DB configuration
builder.Services.AddDbContext<RecipesDbContext>(options =>
{
    options.UseNpgsql(builder.Configuration.GetConnectionString("DefaultConnection"));
});

var app = builder.Build();

// Enable authentication and authorization in the middleware pipeline
app.UseAuthentication();
app.UseAuthorization();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

var api = app.MapGroup("/api");

api.MapAuthRoutes();
api.MapUserRoutes();
api.MapRecipeRoutes();
api.MapIngredientRoutes();

app.Run();

