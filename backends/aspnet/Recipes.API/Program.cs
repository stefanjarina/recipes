using System.Text.Json;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Http.Json;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;
using Recipes.API.auth;
using Recipes.API.db;
using Recipes.API.handlers;
using Recipes.API.services;

var builder = WebApplication.CreateBuilder(args);

// Json configuration
builder.Services.Configure<JsonOptions>(options =>
{
    options.SerializerOptions.PropertyNamingPolicy = JsonNamingPolicy.SnakeCaseLower;
});

// Load JWT settings from appsettings.json
var jwtSettingsSection = builder.Configuration.GetSection("JwtSettings");
builder.Services.Configure<JwtSettings>(jwtSettingsSection);

var jwtSettings = jwtSettingsSection.Get<JwtSettings>();
var key = System.Text.Encoding.ASCII.GetBytes(jwtSettings.SecretKey);

// Add authentication services
builder.Services.AddAuthentication(options =>
{
    options.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme;
    options.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme;
})
.AddJwtBearer(options =>
{
    options.RequireHttpsMetadata = false; // Set to true in production
    options.SaveToken = true;
    options.TokenValidationParameters = new TokenValidationParameters
    {
        ValidateIssuerSigningKey = true,
        IssuerSigningKey = new SymmetricSecurityKey(key),
        ValidateIssuer = true,
        ValidateAudience = false,
        ValidIssuer = jwtSettings.Issuer,
        ValidAudience = jwtSettings.Audience,
        ClockSkew = TimeSpan.Zero // Optional: to ensure the token expires exactly at the token expiration time
    };
});

// Add authorization services
builder.Services.AddAuthorization();

// DB configuration
builder.Services.AddDbContext<RecipesDbContext>(options =>
{
    options.UseNpgsql(builder.Configuration.GetConnectionString("DefaultConnection"));
});

// API configuration
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// USER SERVICES
builder.Services.AddScoped<ITokenService, TokenService>();

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

api.MapPost("/login", AuthHandlers.LoginAsync).WithOpenApi();

api.MapGet("/users", UserHandlers.GetAllUsers).RequireAuthorization().WithOpenApi();
api.MapGet("/users/{id:guid}", UserHandlers.GetUserById).WithOpenApi();

app.Run();

