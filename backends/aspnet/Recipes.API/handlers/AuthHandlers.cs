using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Recipes.API.db;
using Recipes.API.dtos;
using Recipes.API.helpers;
using Recipes.API.services;

namespace Recipes.API.handlers;

public static class AuthHandlers
{
    public static async Task<IResult> LoginAsync([FromBody] LoginDto loginDto,
        IConfiguration configuration, ITokenService tokenService, RecipesDbContext db)
    {
        var user = await db.Users.FirstOrDefaultAsync(u => u.Email == loginDto.Email);
        if (user == null)
            return ApiResponse.NotFound();

        if (!BCrypt.Net.BCrypt.EnhancedVerify(loginDto.Password, user.PasswordHash))
            return ApiResponse.Unauthorized();

        var token = tokenService.GenerateToken(user);

        var response = new
        {
            full_name = user.FullName,
            email = user.Email,
            token,
        };
        
        return ApiResponse.Ok(response);
    }
}