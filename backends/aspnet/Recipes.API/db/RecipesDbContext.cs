using System;
using System.Collections.Generic;
using Microsoft.EntityFrameworkCore;
using Recipes.API.entities;

namespace Recipes.API.db;

public partial class RecipesDbContext : DbContext
{
    public RecipesDbContext()
    {
    }

    public RecipesDbContext(DbContextOptions<RecipesDbContext> options)
        : base(options)
    {
    }

    public virtual DbSet<Comment> Comments { get; set; }

    public virtual DbSet<Ingredient> Ingredients { get; set; }

    public virtual DbSet<Measurement> Measurements { get; set; }

    public virtual DbSet<MeasurementType> MeasurementTypes { get; set; }

    public virtual DbSet<Photo> Photos { get; set; }

    public virtual DbSet<Rating> Ratings { get; set; }

    public virtual DbSet<Recipe> Recipes { get; set; }

    public virtual DbSet<Step> Steps { get; set; }

    public virtual DbSet<StepIngredient> StepIngredients { get; set; }

    public virtual DbSet<User> Users { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        => optionsBuilder.UseNpgsql("Name=ConnectionStrings:DefaultConnection");

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<Comment>(entity =>
        {
            entity.HasKey(e => e.CommentId).HasName("comments_pkey");

            entity.ToTable("comments");

            entity.HasIndex(e => e.RecipeId, "idx_comment_recipe");

            entity.HasIndex(e => e.UserId, "idx_comment_user");

            entity.Property(e => e.CommentId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("comment_id");
            entity.Property(e => e.CommentText).HasColumnName("comment_text");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.RecipeId).HasColumnName("recipe_id");
            entity.Property(e => e.UpdatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("updated_at");
            entity.Property(e => e.UserId).HasColumnName("user_id");

            entity.HasOne(d => d.Recipe).WithMany(p => p.Comments)
                .HasForeignKey(d => d.RecipeId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("comments_recipe_id_fkey");

            entity.HasOne(d => d.User).WithMany(p => p.Comments)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("comments_user_id_fkey");
        });

        modelBuilder.Entity<Ingredient>(entity =>
        {
            entity.HasKey(e => e.IngredientId).HasName("ingredients_pkey");

            entity.ToTable("ingredients");

            entity.HasIndex(e => e.RecipeId, "idx_ingredient_recipe");

            entity.Property(e => e.IngredientId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("ingredient_id");
            entity.Property(e => e.Amount).HasColumnName("amount");
            entity.Property(e => e.AmountMetric).HasColumnName("amount_metric");
            entity.Property(e => e.MeasurementId).HasColumnName("measurement_id");
            entity.Property(e => e.Name)
                .HasMaxLength(255)
                .HasColumnName("name");
            entity.Property(e => e.RecipeId).HasColumnName("recipe_id");

            entity.HasOne(d => d.Measurement).WithMany(p => p.Ingredients)
                .HasForeignKey(d => d.MeasurementId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("ingredients_measurement_id_fkey");

            entity.HasOne(d => d.Recipe).WithMany(p => p.Ingredients)
                .HasForeignKey(d => d.RecipeId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("ingredients_recipe_id_fkey");
        });

        modelBuilder.Entity<Measurement>(entity =>
        {
            entity.HasKey(e => e.MeasurementId).HasName("measurements_pkey");

            entity.ToTable("measurements");

            entity.HasIndex(e => e.Name, "measurements_name_key").IsUnique();

            entity.Property(e => e.MeasurementId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("measurement_id");
            entity.Property(e => e.Name)
                .HasMaxLength(100)
                .HasColumnName("name");
            entity.Property(e => e.System)
                .HasMaxLength(10)
                .HasColumnName("system");
            entity.Property(e => e.TypeId).HasColumnName("type_id");

            entity.HasOne(d => d.Type).WithMany(p => p.Measurements)
                .HasForeignKey(d => d.TypeId)
                .HasConstraintName("measurements_type_id_fkey");
        });

        modelBuilder.Entity<MeasurementType>(entity =>
        {
            entity.HasKey(e => e.TypeId).HasName("measurement_types_pkey");

            entity.ToTable("measurement_types");

            entity.HasIndex(e => e.Name, "measurement_types_name_key").IsUnique();

            entity.Property(e => e.TypeId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("type_id");
            entity.Property(e => e.Name)
                .HasMaxLength(50)
                .HasColumnName("name");
        });

        modelBuilder.Entity<Photo>(entity =>
        {
            entity.HasKey(e => e.PhotoId).HasName("photos_pkey");

            entity.ToTable("photos");

            entity.HasIndex(e => e.RecipeId, "idx_photo_recipe");

            entity.Property(e => e.PhotoId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("photo_id");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.FilePath).HasColumnName("file_path");
            entity.Property(e => e.RecipeId).HasColumnName("recipe_id");
            entity.Property(e => e.UpdatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("updated_at");
            entity.Property(e => e.Url).HasColumnName("url");

            entity.HasOne(d => d.Recipe).WithMany(p => p.Photos)
                .HasForeignKey(d => d.RecipeId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("photos_recipe_id_fkey");
        });

        modelBuilder.Entity<Rating>(entity =>
        {
            entity.HasKey(e => e.RatingId).HasName("ratings_pkey");

            entity.ToTable("ratings");

            entity.Property(e => e.RatingId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("rating_id");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.Rating1).HasColumnName("rating");
            entity.Property(e => e.RecipeId).HasColumnName("recipe_id");
            entity.Property(e => e.UpdatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("updated_at");
            entity.Property(e => e.UserId).HasColumnName("user_id");

            entity.HasOne(d => d.Recipe).WithMany(p => p.Ratings)
                .HasForeignKey(d => d.RecipeId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("ratings_recipe_id_fkey");

            entity.HasOne(d => d.User).WithMany(p => p.Ratings)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("ratings_user_id_fkey");
        });

        modelBuilder.Entity<Recipe>(entity =>
        {
            entity.HasKey(e => e.RecipeId).HasName("recipes_pkey");

            entity.ToTable("recipes");

            entity.HasIndex(e => e.UserId, "idx_recipe_user");

            entity.Property(e => e.RecipeId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("recipe_id");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.Description).HasColumnName("description");
            entity.Property(e => e.Title)
                .HasMaxLength(255)
                .HasColumnName("title");
            entity.Property(e => e.UpdatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("updated_at");
            entity.Property(e => e.UserId).HasColumnName("user_id");
            entity.Property(e => e.Visibility)
                .HasMaxLength(10)
                .HasDefaultValueSql("'public'::character varying")
                .HasColumnName("visibility");

            entity.HasOne(d => d.User).WithMany(p => p.Recipes)
                .HasForeignKey(d => d.UserId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("recipes_user_id_fkey");
        });

        modelBuilder.Entity<Step>(entity =>
        {
            entity.HasKey(e => e.StepId).HasName("steps_pkey");

            entity.ToTable("steps");

            entity.HasIndex(e => e.RecipeId, "idx_step_recipe");

            entity.Property(e => e.StepId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("step_id");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.Instructions).HasColumnName("instructions");
            entity.Property(e => e.RecipeId).HasColumnName("recipe_id");
            entity.Property(e => e.StepNumber).HasColumnName("step_number");
            entity.Property(e => e.Title)
                .HasMaxLength(255)
                .HasColumnName("title");
            entity.Property(e => e.UpdatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("updated_at");

            entity.HasOne(d => d.Recipe).WithMany(p => p.Steps)
                .HasForeignKey(d => d.RecipeId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("steps_recipe_id_fkey");
        });

        modelBuilder.Entity<StepIngredient>(entity =>
        {
            entity.HasKey(e => e.StepIngredientId).HasName("step_ingredients_pkey");

            entity.ToTable("step_ingredients");

            entity.HasIndex(e => e.StepId, "idx_step_ingredient_step");

            entity.Property(e => e.StepIngredientId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("step_ingredient_id");
            entity.Property(e => e.Amount).HasColumnName("amount");
            entity.Property(e => e.AmountMetric).HasColumnName("amount_metric");
            entity.Property(e => e.IngredientId).HasColumnName("ingredient_id");
            entity.Property(e => e.MeasurementId).HasColumnName("measurement_id");
            entity.Property(e => e.StepId).HasColumnName("step_id");

            entity.HasOne(d => d.Ingredient).WithMany(p => p.StepIngredients)
                .HasForeignKey(d => d.IngredientId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("step_ingredients_ingredient_id_fkey");

            entity.HasOne(d => d.Measurement).WithMany(p => p.StepIngredients)
                .HasForeignKey(d => d.MeasurementId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("step_ingredients_measurement_id_fkey");

            entity.HasOne(d => d.Step).WithMany(p => p.StepIngredients)
                .HasForeignKey(d => d.StepId)
                .OnDelete(DeleteBehavior.Cascade)
                .HasConstraintName("step_ingredients_step_id_fkey");
        });

        modelBuilder.Entity<User>(entity =>
        {
            entity.HasKey(e => e.UserId).HasName("users_pkey");

            entity.ToTable("users");

            entity.HasIndex(e => e.Email, "idx_user_email");

            entity.HasIndex(e => e.Email, "users_email_key").IsUnique();

            entity.Property(e => e.UserId)
                .HasDefaultValueSql("gen_random_uuid()")
                .HasColumnName("user_id");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.Email)
                .HasMaxLength(255)
                .HasColumnName("email");
            entity.Property(e => e.FullName)
                .HasMaxLength(255)
                .HasColumnName("full_name");
            entity.Property(e => e.PasswordHash).HasColumnName("password_hash");
            entity.Property(e => e.UpdatedAt)
                .HasDefaultValueSql("CURRENT_TIMESTAMP")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("updated_at");
        });

        OnModelCreatingPartial(modelBuilder);
    }

    partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
}
