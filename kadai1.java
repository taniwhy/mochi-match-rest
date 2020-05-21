

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Calendar;
import java.util.Date;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;


/**
 * Servlet implementation class Kadai1
 */
@WebServlet("/Kadai1")
public class Kadai1 extends HttpServlet {
	private static final long serialVersionUID = 1L;

    /**
     * @see HttpServlet#HttpServlet()
     */
    public Kadai1() {
        super();
        // TODO Auto-generated constructor stub
    }

	/**
	 * @see HttpServlet#doGet(HttpServletRequest request, HttpServletResponse response)
	 */
	protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		String[] messages = { "おはようございます。朝です", "こんにちは。昼です。", "こんばんは。夜です", "夜中です。寝ましょう" };
		int index;
		Date d = new Date();
	    Calendar c = Calendar.getInstance();
	    c.setTime(d);
	    int hour = c.get(Calendar.HOUR_OF_DAY);

	    if(5 <= hour && hour <= 10) {
	    	index = 0;
	    } else if (11 <= hour && hour <= 17) {
	    	index = 1;
	    } else if (18 <= hour && hour <= 22) {
	    	index = 2;
	    } else {
	    	index = 3;
	    }

		//　HTMLを出力
		response.setContentType("text/html; charset=UTF8");
		PrintWriter out = response.getWriter();
		out.println("<html>");
		out.println("<head>");
		out.println("<title>スッキリ占い</title>");
		out.println("</head>");
		out.println("<body>");
		out.println("<p>" + messages[index] + "</p>");
		out.println("</body>");
		out.println("</html>");
	}

}
