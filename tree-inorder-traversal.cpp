

/* you only have to complete the function given below.  
Node is defined as  

class Node {
    public:
        int data;
        Node *left;
        Node *right;
        Node(int d) {
            data = d;
            left = NULL;
            right = NULL;
        }
};
*/

/*
Algorithm Inorder(tree)
   1. Traverse the left subtree, i.e., call Inorder(left-subtree)
   2. Visit the root.
   3. Traverse the right subtree, i.e., call Inorder(right-subtree)
*/
    void inOrder(Node *root) {
        if (root->left != nullptr) {
            inOrder(root->left);
        }
        std::cout << root->data << " ";
        if (root->right != nullptr) {
            inOrder(root->right);
        }
    }

