using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Recipes.API.db;
using Recipes.API.dtos;
using Recipes.API.services;

namespace Recipes.API.handlers;

public static class AuthHandlers
{
    public static async Task<IResult> LoginAsync([FromBody] LoginDto loginDto,
        IConfiguration configuration, ITokenService tokenService, RecipesDbContext db)
    {
        var user = await db.Users.FirstOrDefaultAsync(u => u.Email == loginDto.Email);
        if (user == null)
            return Results.NotFound();

        if (user.PasswordHash != loginDto.Password)
            return Results.Unauthorized();

        var token = tokenService.GenerateToken(user);

        var response = new
        {
            full_name = user.FullName,
            email = user.Email,
            token,
        };
        
        return Results.Json(response);
    }
}