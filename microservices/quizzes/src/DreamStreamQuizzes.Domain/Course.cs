namespace DreamStreamQuizzes.Domain;

[Table("courses", Schema = "quizzes")]
public class Course
{
    [Column("id", TypeName = "uuid")]
    [Key]
    public Guid CourseId { get; set; }

    [Column("category_id", TypeName = "uuid")]
    [ForeignKey("categories")]
    public Guid CategoryId { get; set; }

    [Column("created_at", TypeName = "timestamp with time zone")]
    public DateTimeOffset CreatedAt { get; set; }

    [Column("edited_at", TypeName = "timestamp with time zone")]
    public DateTimeOffset EditedAt { get; set; }

    [Column("name", TypeName = "text")]
    public String Name { get; set; }
}