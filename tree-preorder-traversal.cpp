

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

Algorithm Preorder(tree)
   1. Visit the root.
   2. Traverse the left subtree, i.e., call Preorder(left-subtree)
   3. Traverse the right subtree, i.e., call Preorder(right-subtree

*/

/*
   Preorder
   1. Visit the root.
   2. Traverse the left subtree, i.e., call Preorder(left-subtree)
   3. Traverse the right subtree, i.e., call Preorder(right-subtree
*/
    void preOrder(Node *root) {
        Node *current;
        std::cout << root->data << " ";
        if (root->left != nullptr) {
            preOrder(root->left);
        }
        if (root->right != nullptr) {
            preOrder(root->right);
        }
    }

