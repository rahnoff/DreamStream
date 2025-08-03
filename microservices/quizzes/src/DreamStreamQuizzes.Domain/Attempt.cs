namespace DreamStreamQuizzes.Domain;

[Table("attempts", Schema = "quizzes")]
public class Attempt
{
    [Column("id", TypeName = "uuid")]
    [Key]
    public Guid AttemptId { get; set; }

    [Column("created_at", TypeName = "timestamp with time zone")]
    public DateTimeOffset CreatedAt { get; set; }

    [Column("edited_at", TypeName = "timestamp with time zone")]
    public DateTimeOffset EditedAt { get; set; }

    [Column("name", TypeName = "text")]
    public String Name { get; set; }
}