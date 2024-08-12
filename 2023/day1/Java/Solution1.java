import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

public class Solution1 {
   public static void main(String[] args) {
      try {
         File input = new File("puzzle_input.txt");
         FileReader fileReader = new FileReader(input);
         BufferedReader reader = new BufferedReader(fileReader);

         String line;
         int result = 0;
         while ((line = reader.readLine()) != null) {
            ArrayList<Integer> arr = new ArrayList<>();
            for (int i = 0; i < line.length(); i++) {
               char c = line.charAt(i);
               if (c <= '9' && c >= '0') {
                  arr.add(c - '0');
               }
            }
            int num = arr.get(0);
            num = num * 10 + arr.get(arr.size() - 1);
            result += num;
         }

         System.out.println("Ans is:" + result);

         reader.close();
      } catch (IOException e) {
         e.printStackTrace();
      }
   }
}
