# First, ensure that the directory containing libtensorflow.so.2 
# is in your LD_LIBRARY_PATH. The LD_LIBRARY_PATH environment variable tells the dynamic 
# linker where to look for shared libraries. You can check its current value with
echo $LD_LIBRARY_PATH

# If /usr/local/lib (where libtensorflow.so.2 is located) is not included, 
# you can add it temporarily for testing purposes:
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

# If you want to make this change permanent, you can add the following line to your .bashrc file:
echo 'export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH' >> ~/.bashrc

# Now, you can run the following command to check if the library is found:
ldconfig -p | grep libtensorflow