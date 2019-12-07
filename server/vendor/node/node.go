//  Author: Shaquille Hall
//  Date: December 6, 2019
//  Title: Blockchain Node Structure in Golang

package OpenFacesNode

type Node struct {
	Idx int   
	Data string
  Hash string
  HashOfPrev string
  IsMined bool
  IsValid bool
  Nonce int
}


func ConstructNode(idx int, info string, hash string, nonce int, prevHash string, isValid bool, isMined bool) (Node) {
	return Node{
			Idx: idx,
  		Data: info,
  		Hash: hash,
  		Nonce: nonce,
  		HashOfPrev: prevHash,
  		IsValid: isValid,
  		IsMined: isMined,
 	}
}

