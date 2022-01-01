
/*The tree node has data, left child and right child 
class Node {
    int data;
    Node* left;
    Node* right;
};

*/
    int height(Node* root) {
        int leftHeight;
        leftHeight = 0;
        if (root->left != nullptr) {
            leftHeight = 1;
            traversal(root->left, &leftHeight);
        }
        int rightHeight;
        rightHeight = 0;
        if (root->right != nullptr) {
          rightHeight = 1;
          traversal(root->right, &rightHeight);
        }  
          if (rightHeight > leftHeight) {
            return rightHeight;
        }
        return leftHeight;
    }

    void traversal(Node *root, int *counter) {
      if (root->right != nullptr) {
        (*counter)++;
        traversal(root->right, counter);
      } else if (root->left != nullptr) {
        (*counter)++;
        traversal(root->left, counter);
      }
    }

