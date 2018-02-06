import java.io.*;

import java.security.InvalidKeyException;
import java.security.KeyFactory;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.spec.InvalidKeySpecException;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;

import javax.crypto.BadPaddingException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;

class test  
{
    private static final String XCHANNEL_RSA_PUBKEY_STR = "30820122300d06092a864886f70d01010105000382010f003082010a0282010100cbc0dbbfca6ba49ff490a86519e258a33a36b7ad041bc2f4e7ad60d77476f4bbcc479872bc6665189a669fae7e038c475c89c96223ee2b3aca52702978b67ebf0d47a1c38066d58ad33debdbab80f2441838a5fd42dfc36c276fb05f7017c8112b4a6dc7d33a475dcdbc1c6f5500297dcf6307d9d3ce984e80ce0988462f955f052e183dc99a2f31877a5b07f76a85caf3dd554314ebfe4418a2ad58a7eebc8ba4f89cdabb3560be50d563407c4d40c40ccd4ea55825fccf2837c87f3c38042451721729f43e364bd843573168ca153e963bdbe795b612b3a4c7b21e4067edc8ae9a4e6a80b4c42407d66697b87f8785663fd173c63c263d884b99153e3532c10203010001";
    private PublicKey xchannelRsaPubKey;
    
    private static final byte[] toBytes(String s) {
        byte[] bytes;
        bytes = new byte[s.length() / 2];
        for (int i = 0; i < bytes.length; i++) {
            bytes[i] = (byte) Integer.parseInt(s.substring(2 * i, 2 * i + 2),
                    16);
        }
        return bytes;
    }
    
    public static PublicKey toPublicKey(String key) throws NoSuchAlgorithmException, InvalidKeySpecException {
        X509EncodedKeySpec publicKeySpec = new X509EncodedKeySpec(toBytes(key));
        KeyFactory kf = KeyFactory.getInstance("RSA");
        PublicKey publicKey = kf.generatePublic(publicKeySpec);
        return publicKey;
    }
    
    public static byte[] encrypt(PublicKey key, byte[] data)
            throws NoSuchAlgorithmException, NoSuchPaddingException, InvalidKeyException, IllegalBlockSizeException, BadPaddingException {
        Cipher pkCipher = Cipher.getInstance(RSA_NONE_OAEPWITHSHA1ANDMGF1PADDING);
        pkCipher.init(Cipher.ENCRYPT_MODE, key);
        return pkCipher.doFinal(data);
    }
    
    private static final String RSA_NONE_OAEPWITHSHA1ANDMGF1PADDING = "RSA/NONE/OAEPWithSHA1AndMGF1Padding";
    
	public static void main (String[] args) throws java.lang.Exception
	{
	    PublicKey pk = toPublicKey(XCHANNEL_RSA_PUBKEY_STR);
	    encrypt(pk, new byte[]{1,2,3});
		System.out.println("hi "+pk);
	}
}

