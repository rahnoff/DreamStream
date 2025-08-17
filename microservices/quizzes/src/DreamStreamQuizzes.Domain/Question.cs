namespace DreamStreamQuizzes.Domain;

[Table("questions", Schema = "quizzes")]
public class Question
{
    [Column("id", TypeName = "uuid")]
    [Key]
    public Guid QuestionId { get; set; }

    [Column("content", TypeName = "text")]
    public string Content { get; set; }

    [Column("created_at", TypeName = "timestamptz")]
    public DateTimeOffset CreatedAt { get; set; }

    [Column("edited_at", TypeName = "timestamptz")]
    public DateTimeOffset EditedAt { get; set; }

    [Column("quiz_id", TypeName = "uuid")]
    [ForeignKey("quizzes")]
    public Guid QuizId { get; set; }
}