#include "fpga_version.h"


int main()
{
    
    FILE *fin;
    fin = fopen("vocabulary.txt", "r");
    int i,j;

    char** dictionary = new char*[8800];
    for(i =0 ; i< 8800; i++){
	dictionary[i] = new char[50];
    }
    
    for(i = 0; i < 8800; i++){ //clear dictionary
      for(j = 0; j< 50; j++){
        dictionary[i][j]=0; 
      }
    }
    i =0;
    while (fgets(dictionary[i],50,fin)!=NULL){
        i++;
        if(i >= 8800)
            break;
    }
    fclose(fin);

    auto t_start = std::chrono::high_resolution_clock::now();    

    
    Predictor* pred = Newfpga("reorderdata_for_maxDSP_small_diffQ.bin");

    auto t_2 = std::chrono::high_resolution_clock::now();
    
    int * idx = Predictfpga(pred,"./image/clocktower.bin");

    auto t_3 = std::chrono::high_resolution_clock::now();
   
   printf("clocktower: \n");
   for(int j=0;j<15;j++)
   {
       printf("%s ",dictionary[idx[j]-1]);
   }

   idx = Predictfpga(pred,"./image/broccoli.bin");
   printf("broccoli: \n");
   for(int j=0;j<15;j++)
   {
       printf("%d %s ",idx[j],dictionary[idx[j]-1]);
   }
   delete idx;
   for(i =0 ; i< 8800; i++){
	delete dictionary[i];
    }
    delete dictionary;
    
    std::chrono::duration<double> data_prepare = t_2 - t_start; 
    std::chrono::duration<double> kernel_exec = t_3 - t_2; 
    std::cout << "Preprocessing: " << data_prepare.count() << "s" << std::endl;
    std::cout << "Kernel exec time: " << kernel_exec.count() << "s" << std::endl;

    return 0;
}
