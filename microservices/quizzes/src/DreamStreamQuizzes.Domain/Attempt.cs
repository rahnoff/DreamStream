namespace DreamStreamQuizzes.Domain;

[Table("attempts", Schema = "quizzes")]
public class Attempt
{
    [Column("id", TypeName = "uuid")]
    [Key]
    public Guid AttemptId { get; set; }

    [Column("created_at", TypeName = "timestamptz")]
    public DateTimeOffset CreatedAt { get; set; }

    [Column("edited_at", TypeName = "timestamptz")]
    public DateTimeOffset EditedAt { get; set; }

    [Column("name", TypeName = "text")]
<<<<<<< HEAD
    public string Name { get; set; }
=======
    public String Name { get; set; }
>>>>>>> 45a50ea570650c68c34638bc618e2e11f7ea3032
}