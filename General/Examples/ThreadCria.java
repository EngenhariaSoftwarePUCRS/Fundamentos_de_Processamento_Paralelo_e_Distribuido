/* PUCRS - Programacao Concorrente - Fernando Dotti */

class Contadora extends Thread {
    private int id;
    private int n;
    private String s;
	
    public Contadora(int _id, int _n, String _s){
    	id = _id;
		n = _n;
		s = _s;
    }
    public void run() {
       for (int i = 0; i < n; i++) {
	       System.out.println(s + " id: "+id+" -> "+i);   
           i++;
      }
    }
}

class TesteCria {
    public static void main(String[] args) {
	  int max = 1000000;
      Contadora p = new Contadora(1,max," ");
      Contadora q = new Contadora(2,max,"                    ");
      Contadora r = new Contadora(3,max,"                                        ");
      Contadora s = new Contadora(4,max,"                                                            ");
	  
	  p.start();
      q.start();
	  
      r.start();
      s.start();
   
      System.out.println("Fim");
    }
}

// Execute o programa e observe o andamento.
// O que se pode supor sobre o avanço relativo das threads ?
