class Useless {
    public static void main(String[] args) {
        String a = args.length > 0 ? args[0] : "123";
        
        String[] aArr = new String[a.length()];
        for (int j = 0, i = a.length() - 1; i >= 0; i--, j++)
            aArr[j] = String.valueOf(a.charAt(i));

        int result = 0;
        for (int i = 0; i < a.length(); i++)
            result += Integer.parseInt(aArr[i]) * (int) Math.pow(10, i);

        System.out.println(result);
    }
}
