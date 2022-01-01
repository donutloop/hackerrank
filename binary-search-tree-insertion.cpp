

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

      Binary Search Tree, is a node-based binary tree data structure which has
      the following properties: The left subtree of a node contains only nodes
      with keys lesser than the node’s key. The right subtree of a node contains
      only nodes with keys greater than the node’s key. The left and right
      subtree each must also be a binary search tree. There must be no duplicate
      nodes.
      */

      Node *insert(Node *root, int data) {
        if (root == nullptr) {
          root = new Node(data);
          return root;
        }
        insertion(root, data);
        return root; 
      }

      void insertion(Node *node, int data) {
        if (node == nullptr) {return;}
        if ((*node).data < data) {
          if (node->right == nullptr) {
            Node *n = new Node(data);
            node->right = n;
          } else {
            insertion(node->right, data);
          }
          return;
          } else if ((*node).data > data) {
            if (node->left == nullptr) {
              Node *n = new Node(data);
              node->left = n;
            } else {
              insertion(node->left, data);
            }
            return;
          }
      }

