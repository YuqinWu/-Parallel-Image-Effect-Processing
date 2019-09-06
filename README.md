# Parallel-Image-Effect-Processing
Implemented an image effect, such as edge detection, blurring, and sharpening, processing program incorporating with multi-threaded technique.


In this project, I basically implemented a concurrent images effects application. This application can read in a png image file, and apply 4 different effects to the image based on the input parameter. The four effects are grayscale, blur, edge detection, and sharpen. User can use the capital letter to indicate which effect to apply. They will be applied in the given order, and store the result image into user specified address.


In the application, the user input is simulated as a csv file with images. The format of the csv file is: imageAddress,storeAddress,[effects...]
Usage of the script:
./editor [-p=[num of threads]] <csv file>

You can find test images and csv files here:
    https://www.dropbox.com/s/cxyvd9f6qr5nep4/csvs.zip?dl=0
