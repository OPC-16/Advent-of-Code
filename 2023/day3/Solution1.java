import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Solution1 {

   public static void main(String[] args) {
      Path filePath = Paths.get("test_input.txt");
      Map<Integer, Boolean> map = new HashMap<>();

      try {
         List<String> lines = Files.readAllLines(filePath);

         int sizeOfLines = lines.size();
         int sizeOfLine = lines.get(0).length();
         for (int i = 0; i < sizeOfLines; i++) {
            String currLine = lines.get(i);
            for (int j = 0; j < sizeOfLine; j++) {
               System.out.print(lines.get(i).charAt(j));
            }
         }
      } catch (IOException e) {
         e.printStackTrace();
      }
   }
}
