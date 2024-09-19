using System.Text.Json;
using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Http.Json;
using Microsoft.IdentityModel.Tokens;
using Recipes.API.helpers;
using Recipes.API.models;
using Recipes.API.services;

namespace Recipes.API.extensions;

public static class ApplicationServicesExtensions
{
    public static IServiceCollection AddApplicationServices(this IServiceCollection services)
    {
        // APP SERVICES
        services.AddScoped<ITokenService, TokenService>();
        
        // API configuration
        services.AddEndpointsApiExplorer();
        services.AddSwaggerGen();
        
        // AutoMapper configuration
        services.AddAutoMapper(typeof(MappingProfile));
        
        // Json configuration
        services.Configure<JsonOptions>(options =>
        {
            options.SerializerOptions.PropertyNamingPolicy = JsonNamingPolicy.SnakeCaseLower;
            options.SerializerOptions.DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull;
        });
        
        return services;
    }
    
    public static IServiceCollection AddAuthorizationServices(this IServiceCollection services, IConfiguration config)
    {
        // Load JWT settings from appsettings.json
        var jwtSettingsSection = config.GetSection("JwtSettings");
        services.Configure<JwtSettings>(jwtSettingsSection);

        var jwtSettings = jwtSettingsSection.Get<JwtSettings>();
        var key = System.Text.Encoding.ASCII.GetBytes(jwtSettings!.SecretKey);

        // Add authentication services
        services.AddAuthentication(options =>
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
        services.AddAuthorization();

        // access JWT claims
        services.AddHttpContextAccessor();

        return services;
    }
}