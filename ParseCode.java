import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;

class ParseCode {
	ParseCode() {
	}

	public static void main(String[] args) {
		System.out.print("hello world!");
		getFile();
	}

	// read file
	private static void getFile() {
		try {
			FileInputStream fis = new FileInputStream("./code.txt");
			byte[] buf = new byte[1024];
			int len = -1;
			StringBuilder sb = new StringBuilder();
			while ((len = fis.read(buf)) > 0) {
				String s = new String(buf, 0, len);
				sb.append(s);
			}
			String[] songArray = sb.toString().split("OxO");
			System.out.println(songArray[0]);
			fis.close();
		} catch (FileNotFoundException e) {
			e.printStackTrace();
		} catch (IOException e) {
			e.printStackTrace();
		}
	
	}

}
