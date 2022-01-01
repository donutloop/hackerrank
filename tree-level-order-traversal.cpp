
/*
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
        void levelOrder(Node *root) {
          int h;
          h = height(root);
          std::cout << root->data << " ";
          for (int i = 0; i < h; i++) {
            printLevel(root, i);
          }
        }
        void printLevel(Node *root, int level) {
          int counter;
          counter = 0;
          print(root, level, counter);
        }
        void print(Node *root, int level, int counter) {
          if (root->left != nullptr && (level == counter)) {
            std::cout << root->left->data << " ";
          }
          if (root->right != nullptr && (level == counter)) {
            std::cout << root->right->data << " ";
          }
          if (counter == level) {
              return;
          }
          counter++;
          if (root->left != nullptr) {
            print(root->left, level, counter);
          }
          if (root->right != nullptr) {
            print(root->right, level, counter);
          }
        }
        int height(Node *root) {
          int leftHeight;
          leftHeight = 0;
          if (root->left != nullptr) {
            leftHeight = 1;
            leftHeight = traversal(root->left, leftHeight);
          }
          int rightHeight;
          rightHeight = 0;
          if (root->right != nullptr) {
            rightHeight = 1;
           rightHeight = traversal(root->right, rightHeight);
          }
          if (rightHeight > leftHeight) {
            return rightHeight;
          }
          return leftHeight;
        }

        int traversal(Node *root, int counter) {
          int rightHeight;
          if (root->right != nullptr) {
            rightHeight = traversal(root->right, counter + 1);
          } else {
            rightHeight = counter;
          }
          int leftHeight;
          if (root->left != nullptr) {
            leftHeight = traversal(root->left, counter + 1);
          } else {
            leftHeight = counter;
          }
          if (leftHeight > rightHeight) {
              return leftHeight;
          }
          return rightHeight;
        }

