namespace DreamStreamQuizzes.Domain;

[Table("authors", Schema = "quizzes")]
public class Author
{
    [Column("id", TypeName = "uuid")]
    [Key]
    public Guid AuthorId { get; set; }
}