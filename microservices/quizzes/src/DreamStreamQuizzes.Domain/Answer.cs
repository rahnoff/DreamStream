namespace DreamStreamQuizzes.Domain;

[Table("answers", Schema = "quizzes")]
public class Answer
{
    [Column("id", TypeName = "uuid")]
    [Key]
    public Guid AnswerId { get; set; }

    [Column("content", TypeName = "text")]
    public string Content { get; set; }

    [Column("correctness", TypeName = "bool")]
    public bool Correctness { get; set; }

    [Column("created_at", TypeName = "timestamptz")]
    public DateTimeOffset CreatedAt { get; set; }

    [Column("edited_at", TypeName = "timestamptz")]
    public DateTimeOffset EditedAt { get; set; }

    [Column("question_id", TypeName = "uuid")]
    [ForeignKey("questions")]
    public Guid QuestionId { get; set; }
}