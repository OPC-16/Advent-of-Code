import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class Solution1 {
   public static void main(String[] args) {
      String filePath = "test_input.txt";
      final int RED_CUBES = 12;
      final int GREEN_CUBES = 13;
      final int BLUE_CUBES = 14;
      int result = 0;

      try (BufferedReader reader = new BufferedReader(new FileReader(filePath))) {
         String line;
         while ((line = reader.readLine()) != null) {

            // getting the game id
            int gameID = line.charAt(5) - '0';
            if (line.charAt(6) != ':') {
               gameID = gameID * 10 + (line.charAt(6) - '0');

               if (line.charAt(7) != ':') {
                  gameID = gameID * 10 + (line.charAt(7) - '0');
               }
            }

            //slicing the string from ":"
            int index = line.indexOf(":");
            String allCubes = line.substring(index + 2);

            int num = 0;
            for (int i = 0; i < allCubes.length(); i++) {
               char c = allCubes.charAt(i);
               
               if (c <= '9' && c >= '0') {
                  num = num * 10 + (c - '0');
                  continue;
               }
               else if (c == ' ') {
                  continue;
               }

               if (c == 'g') {
                  if (num > GREEN_CUBES) {
                     i = Math.min(allCubes.indexOf(","), allCubes.indexOf(";")) + 1;
                  }
                  else {
                     result += gameID;
                  }
               }
               else if (c == 'r') {
                  if (num > RED_CUBES) {
                     i = Math.min(allCubes.indexOf(","), allCubes.indexOf(";")) + 1;
                  }
                  else {
                     result += gameID;
                  }
               }
               else if (c == 'b') {
                  if (num > BLUE_CUBES) {
                     i = Math.min(allCubes.indexOf(","), allCubes.indexOf(";")) + 1;
                  }
                  else {
                     result += gameID;
                  }
               }
               num = 0;
            }

         }//while
         System.out.println(result);

      } catch (IOException e) {
         e.printStackTrace();
      }
   }
}
