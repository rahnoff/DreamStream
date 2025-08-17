namespace DreamStreamQuizzes.Domain;

[Table("quizzes", Schema = "quizzes")]
public class Quiz
{
    [Column("id", TypeName = "uuid")]
    [Key]
    public Guid QuizId { get; set; }

    [Column("course_id", TypeName = "uuid")]
    [ForeignKey("courses")]
    public Guid CourseId { get; set; }

    [Column("created_at", TypeName = "timestamptz")]
    public DateTimeOffset CreatedAt { get; set; }

    [Column("edited_at", TypeName = "timestamptz")]
    public DateTimeOffset EditedAt { get; set; }
}