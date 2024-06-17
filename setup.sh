# First, ensure that the directory containing libtensorflow.so.2 
# is in your LD_LIBRARY_PATH. The LD_LIBRARY_PATH environment variable tells the dynamic 
# linker where to look for shared libraries. You can check its current value with
echo $LD_LIBRARY_PATH

# If /usr/local/lib (where libtensorflow.so.2 is located) is not included, 
# you can add it temporarily for testing purposes:
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

# Anyway ^, this command will only add the directory to LD_LIBRARY_PATH for the current shell session.
# Although the change will be lost when you close the terminal, but it will cache the library path on
# `/etc/ld.so.cache` file.

# If you want to make this change permanent, you can add the following line to your .bashrc file:
echo 'export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH' >> ~/.bashrc

# Now, you can run the following command to check if the library is found:
ldconfig -p | grep libtensorflow

# Notes:
# LD_LIBRARY_PATH is an environment variable that the system's dynamic linker (ld.so) 
# uses to locate shared libraries (.so files) that are required by a program at runtime. 
# When we  run a program that depends on shared libraries, the dynamic linker searches for those 
# libraries in a predefined set of directories (/lib, /usr/lib, etc.) 
# and also in any directories specified in LD_LIBRARY_PATH.