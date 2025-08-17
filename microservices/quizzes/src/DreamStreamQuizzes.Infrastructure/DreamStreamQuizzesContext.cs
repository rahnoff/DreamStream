namespace DreamStreamQuizzes.Domain;

public class DreamStreamQuizzesContext : DbContext
{
    public DreamStreamQuizzesContext(DbContextOptions<DreamStreamQuizzesContext> options) : base(options)
    {

    }

    public DbSet<Answer> Answers { get; set; }
    public DbSet<Attempt> Attempts { get; set; }
    public DbSet<Author> Authors { get; set; }
    public DbSet<Category> Categories { get; set; }
    public DbSet<Course> Courses { get; set; }
    public DbSet<Employee> Employees { get; set; }
    public DbSet<Question> Questions { get; set; }
    public DbSet<Quiz> Quizzes { get; set; }
}